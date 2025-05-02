package middleware

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/pkg"
	"github.com/gofiber/fiber/v3"
)

type serviceToken struct {
	apiAuthService interfaces.ApiAuthInterface
}

func NewApiTokenMiddleware(apiAuthService interfaces.ApiAuthInterface) interfaces.ApiAuthMiddleware {
	return &serviceToken{
		apiAuthService: apiAuthService,
	}
}

func (s serviceToken) ApiAuthValidator(c fiber.Ctx) error {
	var inputBody dto.ReqHeaderEntry
	errResponse, err := pkg.ValidateRequestBody(&inputBody, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errResponse)
	}

	userID := c.Locals("userID").(uint)
	tokens := s.apiAuthService.GetClientAll(userID, inputBody.ClientID)
	if len(tokens) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: "ClientID has not any tokens",
		})
	}

	return c.Next()
}
