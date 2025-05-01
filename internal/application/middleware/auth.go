package middleware

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/gofiber/fiber/v3"
)

type authMiddleware struct {
	tokenService interfaces.TokenService
}

func NewAuthTokenMiddleware(tokenService interfaces.TokenService) interfaces.TokenMiddleware {
	return &authMiddleware{tokenService: tokenService}
}

func (a *authMiddleware) CheckTokenAuth(c fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{Message: "No token provided"})
	}
	tokenM, err := a.tokenService.GetByToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{Message: "Invalid token"})
	}
	userIP := c.Host()
	if tokenM.AccessIP != "" && tokenM.AccessIP != userIP {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{Message: "Invalid token"})
	}
	c.Locals("userID", tokenM.UserID)
	return c.Next()
}
