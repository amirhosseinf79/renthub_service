package interfaces

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/gofiber/fiber/v3"
)

type ApiAuthInterface interface {
	GetByUnique(userID uint, clientID string, service string) (*models.ApiAuth, error)
	GetClientAll(userID uint, clientID string) []*models.ApiAuth
	UpdateOrCreate(userID uint, fields dto.ApiAuthRequest) error
	SignOutService(userID uint, fields dto.ApiAuthSignOut) error
}

type ApiAuthMiddleware interface {
	ApiAuthValidator(c fiber.Ctx) error
}
