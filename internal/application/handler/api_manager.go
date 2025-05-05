package handler

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/pkg"
	"github.com/gofiber/fiber/v3"
)

type handlerSt struct {
	serviceManagerC interfaces.BrokerClientInterface
	ApiAuthService  interfaces.ApiAuthInterface
	defaultResponse dto.ErrorResponse
}

func NewManagerHandler(
	serviceManagerC interfaces.BrokerClientInterface,
	apiAuthService interfaces.ApiAuthInterface,
) interfaces.ManagerHandlerInterface {
	return &handlerSt{
		serviceManagerC: serviceManagerC,
		ApiAuthService:  apiAuthService,
		defaultResponse: dto.ErrorResponse{Message: "ok"},
	}
}

func (h *handlerSt) UpdatePrice(ctx fiber.Ctx) error {
	var inputBody dto.EditPriceRequest
	ctx.Bind().Body(&inputBody)
	userID := ctx.Locals("userID").(uint)

	taskBody := dto.ClientUpdateBody{
		UserID:   userID,
		Header:   inputBody.ReqHeaderEntry,
		Services: inputBody.Prices,
		Dates:    inputBody.Dates,
	}

	h.serviceManagerC.AsyncUpdate("price", taskBody)
	return ctx.JSON(h.defaultResponse)
}

func (h *handlerSt) UpdateDiscount(ctx fiber.Ctx) error {
	var inputBody dto.EditDiscountRequest
	ctx.Bind().Body(&inputBody)
	userID := ctx.Locals("userID").(uint)

	taskBody := dto.ClientUpdateBody{
		UserID:          userID,
		Header:          inputBody.ReqHeaderEntry,
		Services:        inputBody.Sites,
		Dates:           inputBody.Dates,
		DiscountPercent: inputBody.DiscountPercent,
	}

	h.serviceManagerC.AsyncUpdate("discount", taskBody)
	return ctx.JSON(h.defaultResponse)
}

func (h *handlerSt) UpdateMinNight(ctx fiber.Ctx) error {
	var inputBody dto.EditMinNightRequest
	ctx.Bind().Body(&inputBody)
	userID := ctx.Locals("userID").(uint)

	taskBody := dto.ClientUpdateBody{
		UserID:    userID,
		Header:    inputBody.ReqHeaderEntry,
		Services:  inputBody.Sites,
		Dates:     inputBody.Dates,
		LimitDays: inputBody.LimitDays,
	}

	h.serviceManagerC.AsyncUpdate("minNight", taskBody)
	return ctx.JSON(h.defaultResponse)
}

func (h *handlerSt) UpdateCalendar(ctx fiber.Ctx) error {
	var inputBody dto.EditCalendarRequest
	ctx.Bind().Body(&inputBody)
	userID := ctx.Locals("userID").(uint)

	taskBody := dto.ClientUpdateBody{
		UserID:   userID,
		Header:   inputBody.ReqHeaderEntry,
		Services: inputBody.Sites,
		Dates:    inputBody.Dates,
		Action:   inputBody.Action,
	}

	h.serviceManagerC.AsyncUpdate("calendar", taskBody)
	return ctx.JSON(h.defaultResponse)
}

func (h *handlerSt) RefreshToken(ctx fiber.Ctx) error {
	var inputBody dto.RefreshTokenRequest
	ctx.Bind().Body(&inputBody)
	userID := ctx.Locals("userID").(uint)

	taskBody := dto.ClientUpdateBody{
		UserID:   userID,
		Header:   inputBody.ReqHeaderEntry,
		Services: inputBody.Sites,
	}

	h.serviceManagerC.AsyncUpdate("token", taskBody)
	return ctx.JSON(h.defaultResponse)
}

func (h *handlerSt) CheckAuth(ctx fiber.Ctx) error {
	var inputBody dto.ReqHeaderEntry
	ctx.Bind().Body(&inputBody)
	h.serviceManagerC.AsyncUpdate("checkAuth", dto.ClientUpdateBody{Header: inputBody})
	return ctx.JSON(h.defaultResponse)
}

func (h *handlerSt) SendServiceOTP(ctx fiber.Ctx) error {
	var inputBody dto.OTPSendRequest
	ctx.Bind().Body(&inputBody)
	userID := ctx.Locals("userID").(uint)
	h.serviceManagerC.AsyncOTP("send", dto.OTPBody{
		UserID:      userID,
		ClientID:    inputBody.ClientID,
		Service:     inputBody.Service,
		PhoneNumebr: inputBody.PhoneNumebr,
	})
	return ctx.JSON(h.defaultResponse)
}

func (h *handlerSt) VerifyServiceOTP(ctx fiber.Ctx) error {
	var inputBody dto.OTPVerifyRequest
	ctx.Bind().Body(&inputBody)
	userID := ctx.Locals("userID").(uint)
	h.serviceManagerC.AsyncOTP("verify", dto.OTPBody{
		UserID:      userID,
		ClientID:    inputBody.ClientID,
		Service:     inputBody.Service,
		PhoneNumebr: inputBody.PhoneNumebr,
		Code:        inputBody.Code,
	})
	return ctx.JSON(h.defaultResponse)
}

func (h *handlerSt) TokenLogin(ctx fiber.Ctx) error {
	var inputBody dto.ApiAuthRequest
	response, err := pkg.ValidateRequestBody(&inputBody, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	userID := ctx.Locals("userID").(uint)
	err = h.ApiAuthService.UpdateOrCreate(userID, inputBody)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: err.Error(),
		})
	}
	return ctx.JSON(dto.ErrorResponse{
		Message: "ok",
	})
}
