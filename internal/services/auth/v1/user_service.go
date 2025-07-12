package auth_v1

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/domain/repository"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/pkg"
)

type userService struct {
	userRepo     repository.UserRepository
	tokenService interfaces.TokenService
}

func NewUserService(userRepo repository.UserRepository, tokenService interfaces.TokenService) interfaces.UserService {
	return &userService{
		userRepo:     userRepo,
		tokenService: tokenService,
	}
}

func (s *userService) RegisterUser(creds dto.UserRegister) (*models.Token, error) {
	exists, err := s.userRepo.CheckEmailExists(creds.Email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, dto.ErrEmailExists
	}

	hashedPassword, err := pkg.HashPassword(creds.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Email:       creds.Email,
		Password:    hashedPassword,
		FirstName:   creds.FirstName,
		LastName:    creds.LastName,
		HookToken:   creds.HookToken,
		HookRefresh: creds.HookRefresh,
		RefreshURL:  creds.RefreshURL,
	}

	err = s.userRepo.Create(&user)
	if err != nil {
		return nil, err
	}

	token, err := s.tokenService.GenerateToken(user.ID)
	return token, err
}

func (s *userService) LoginUser(creds dto.UserLogin) (*models.Token, error) {
	user, err := s.userRepo.GetByEmail(creds.Email)
	if err != nil {
		return nil, dto.ErrInvalidCredentials
	}

	if valid := user.ValidatePassword(creds.Password); !valid {
		return nil, dto.ErrInvalidCredentials
	}

	token, err := s.tokenService.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *userService) GetUserById(id uint) (*models.User, error) {
	user, err := s.userRepo.GetByID(id)
	return user, err
}

func (s *userService) UpdateUser(id uint, info dto.UserUpdate) error {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return err
	}

	if info.Email != "" {
		user.Email = info.Email
	}
	if info.Password != "" {
		hashedPassword, err := pkg.HashPassword(info.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	}
	if info.FirstName != "" {
		user.FirstName = info.FirstName
	}
	if info.LastName != "" {
		user.LastName = info.LastName
	}
	if info.HookToken != "" {
		user.HookToken = info.HookToken
	}
	if info.HookRefresh != "" {
		user.HookRefresh = info.HookRefresh
	}
	if info.RefreshURL != "" {
		user.RefreshURL = info.RefreshURL
	}

	err = s.userRepo.Update(user)
	if err != nil {
		return err
	}

	return nil
}
