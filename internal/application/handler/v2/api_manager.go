package handler_v2

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	auth_dto "github.com/amirhosseinf79/renthub_service/internal/dto/auth"
	request_v2 "github.com/amirhosseinf79/renthub_service/internal/dto/request/v2"
	"github.com/amirhosseinf79/renthub_service/pkg"
	"github.com/gofiber/fiber/v3"
)

type handlerSt struct {
	services        map[string]interfaces.ApiService
	serviceManagerC interfaces.BrokerClientInterface_v2
	ApiAuthService  interfaces.ApiAuthInterface
	defaultResponse dto.ErrorResponse
}

func NewManagerHandler(
	services map[string]interfaces.ApiService,
	serviceManagerC interfaces.BrokerClientInterface_v2,
	apiAuthService interfaces.ApiAuthInterface,
) interfaces.ManagerHandlerInterface {
	return &handlerSt{
		services:        services,
		serviceManagerC: serviceManagerC,
		ApiAuthService:  apiAuthService,
		defaultResponse: dto.ErrorResponse{Message: "ok"},
	}
}

func (h *handlerSt) UpdatePrice(ctx fiber.Ctx) error {
	var inputBody request_v2.EditPriceRequest
	ctx.Bind().Body(&inputBody)
	userID := ctx.Locals("userID").(uint)

	taskBody := request_v2.ClientUpdateBody{
		UserID:   userID,
		Header:   inputBody.ReqHeaderEntry,
		Services: inputBody.Prices,
		Dates:    inputBody.Dates,
	}

	h.serviceManagerC.AsyncUpdate("price", taskBody)
	return ctx.JSON(h.defaultResponse)
}

func (h *handlerSt) UpdateDiscount(ctx fiber.Ctx) error {
	var inputBody request_v2.EditDiscountRequest
	ctx.Bind().Body(&inputBody)
	userID := ctx.Locals("userID").(uint)

	taskBody := request_v2.ClientUpdateBody{
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
	var inputBody request_v2.EditMinNightRequest
	ctx.Bind().Body(&inputBody)
	userID := ctx.Locals("userID").(uint)

	taskBody := request_v2.ClientUpdateBody{
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
	var inputBody request_v2.EditCalendarRequest
	ctx.Bind().Body(&inputBody)
	userID := ctx.Locals("userID").(uint)

	taskBody := request_v2.ClientUpdateBody{
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
	var inputBody request_v2.RefreshTokenRequest
	ctx.Bind().Body(&inputBody)
	userID := ctx.Locals("userID").(uint)

	taskBody := request_v2.ClientUpdateBody{
		UserID:   userID,
		Header:   inputBody.ReqHeaderEntry,
		Services: inputBody.Sites,
	}

	h.serviceManagerC.AsyncUpdate("token", taskBody)
	return ctx.JSON(h.defaultResponse)
}

func (h *handlerSt) CheckAuth(ctx fiber.Ctx) error {
	var inputBody request_v2.ReqHeaderWithClientEntry
	ctx.Bind().Body(&inputBody)
	userId := ctx.Locals("userID").(uint)
	h.serviceManagerC.AsyncUpdate("checkAuth",
		request_v2.ClientUpdateBody{Header: request_v2.ReqHeaderEntry{
			UpdateId:    inputBody.UpdateId,
			CallbackUrl: inputBody.CallbackUrl,
			ClientID:    inputBody.ClientID,
		}, UserID: userId})
	return ctx.JSON(h.defaultResponse)
}

func (h *handlerSt) SendServiceOTP(ctx fiber.Ctx) error {
	var inputBody request_v2.OTPSendRequest
	ctx.Bind().Body(&inputBody)
	userID := ctx.Locals("userID").(uint)
	selectedService, ok := h.services[inputBody.Service]
	if !ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: dto.ErrInvalidRequest.Error(),
		})
	}
	model, _ := selectedService.SendOtp(dto.RequiredFields{
		UserID:   userID,
		ClientID: inputBody.ClientID,
	}, inputBody.PhoneNumebr)
	if !model.IsSucceed {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.OTPErrorResponse{
			Message:        dto.ErrInvalidRequest.Error(),
			ServiceMessage: model.FinalResult,
		})
	}
	return ctx.JSON(h.defaultResponse)

	// h.serviceManagerC.AsyncOTP("send", dto.OTPBody{
	// 	UserID:      userID,
	// 	ClientID:    inputBody.ClientID,
	// 	Service:     inputBody.Service,
	// 	PhoneNumebr: inputBody.PhoneNumebr,
	// })
	// return ctx.JSON(h.defaultResponse)
}

func (h *handlerSt) VerifyServiceOTP(ctx fiber.Ctx) error {
	var inputBody request_v2.OTPVerifyRequest
	ctx.Bind().Body(&inputBody)
	userID := ctx.Locals("userID").(uint)
	selectedService, ok := h.services[inputBody.Service]
	if !ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: dto.ErrInvalidRequest.Error(),
		})
	}
	model, _ := selectedService.VerifyOtp(dto.RequiredFields{
		UserID:   userID,
		ClientID: inputBody.ClientID,
	}, dto.OTPCreds{
		PhoneNumber: inputBody.PhoneNumebr,
		OTPCode:     inputBody.Code,
	})
	if !model.IsSucceed {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.OTPErrorResponse{
			Message:        dto.ErrInvalidCode.Error(),
			ServiceMessage: model.FinalResult,
		})
	}
	userToken, err := h.ApiAuthService.GetByUnique(userID, inputBody.ClientID, inputBody.Service)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.OTPErrorResponse{
			Message:        dto.ErrorUnauthorized.Error(),
			ServiceMessage: model.FinalResult,
		})
	}
	response := auth_dto.NewResponse(userToken)
	return ctx.JSON(response)
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

func (h *handlerSt) SignOutClient(ctx fiber.Ctx) error {
	var fields dto.ApiAuthSignOut
	response, err := pkg.ValidateRequestBody(&fields, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	userID := ctx.Locals("userID").(uint)
	err = h.ApiAuthService.SignOutService(userID, fields)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: dto.ErrorUnauthorized.Error(),
		})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
