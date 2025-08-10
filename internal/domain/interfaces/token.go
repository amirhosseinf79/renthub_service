package interfaces

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/gofiber/fiber/v3"
)

type TokenService interface {
	GenerateToken(userID uint, accessIP string) (*models.Token, error)
	RefreshToken(refreshToken string) (*models.Token, error)
	GetByToken(token string) (*models.Token, error)
}

type TokenMiddleware interface {
	CheckTokenAuth(c fiber.Ctx) error
}
