package auth

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
		Email:     creds.Email,
		Password:  hashedPassword,
		FirstName: creds.FirstName,
		LastName:  creds.LastName,
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

func (s *userService) UpdateTokens(id uint, access, refresh string) error {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return err
	}

	user.HookToken = access
	user.HookRefresh = refresh

	err = s.userRepo.Update(user)
	if err != nil {
		return err
	}

	return nil
}
