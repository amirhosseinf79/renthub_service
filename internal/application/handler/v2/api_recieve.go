package handler_v2

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	request_v2 "github.com/amirhosseinf79/renthub_service/internal/dto/request/v2"
	"github.com/gofiber/fiber/v3"
)

type handler struct {
	serviceManagerC interfaces.BrokerClientInterface_v2
	defaultResponse dto.ErrorResponse
	serviceError    dto.ErrorResponse
}

func NewRecieveHandler(
	serviceManagerC interfaces.BrokerClientInterface_v2,
) interfaces.RecieveHandler {
	return &handler{
		serviceManagerC: serviceManagerC,
		defaultResponse: dto.ErrorResponse{Message: "ok"},
		serviceError:    dto.ErrorResponse{Message: dto.ErrService.Error()},
	}
}

func (h *handler) GetReservations(ctx fiber.Ctx) error {
	var fields request_v2.RecieveBody
	ctx.Bind().Body(&fields)
	userID := ctx.Locals("userID").(uint)

	body := request_v2.ClientRecieveBody{
		UserID:   userID,
		Header:   fields.ReqHeaderEntry,
		Services: fields.Sites,
	}

	err := h.serviceManagerC.AsyncRecieve("reservation", body)
	if err != nil {
		return ctx.Status(fiber.StatusServiceUnavailable).JSON(h.serviceError)
	}
	return ctx.JSON(h.defaultResponse)
}
