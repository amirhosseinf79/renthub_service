package interfaces

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/gofiber/fiber/v3"
)

type UserHandler interface {
	RegisterUser(c fiber.Ctx) error
	LoginUser(c fiber.Ctx) error
	UpdateUser(c fiber.Ctx) error
	RefreshToken(c fiber.Ctx) error
}

type UserService interface {
	RegisterUser(creds dto.UserRegister) (*models.Token, error)
	LoginUser(creds dto.UserLogin) (*models.Token, error)
	GetUserById(id uint) (*models.User, error)
	UpdateUser(id uint, info dto.UserUpdate) error
}
