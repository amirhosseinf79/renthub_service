package middleware_v1

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	request_v1 "github.com/amirhosseinf79/renthub_service/internal/dto/request/v1"
	"github.com/amirhosseinf79/renthub_service/pkg"
	"github.com/gofiber/fiber/v3"
)

type serviceToken struct {
	client         interfaces.BrokerClientInterface
	apiAuthService interfaces.ApiAuthInterface
}

func NewApiTokenMiddleware(
	client interfaces.BrokerClientInterface,
	apiAuthService interfaces.ApiAuthInterface,
) interfaces.ApiAuthMiddleware {
	return &serviceToken{
		client:         client,
		apiAuthService: apiAuthService,
	}
}

func (s serviceToken) ApiAuthValidator(c fiber.Ctx) error {
	var inputBody request_v1.ReqHeaderEntry
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

	// taskBody := dto.ClientUpdateBody{
	// 	UserID: userID,
	// 	Header: dto.ReqHeaderEntry{
	// 		CallbackUrl: "ahttp://test/authorization",
	// 		ClientID:    inputBody.ClientID,
	// 	},
	// 	Services: []dto.SiteEntry{},
	// 	Dates:    []string{},
	// }
	// s.client.AsyncUpdate("checkAuth", taskBody)
	return c.Next()
}
