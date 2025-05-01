package handler

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/pkg"
	"github.com/gofiber/fiber/v3"
)

type handlerSt struct {
	serviceManager  interfaces.ServiceManager
	ApiAuthService  interfaces.ApiAuthInterface
	defaultResponse dto.ErrorResponse
}

func NewManagerHandler(
	serviceManager interfaces.ServiceManager,
	apiAuthService interfaces.ApiAuthInterface,
) interfaces.ManagerHandlerInterface {
	return &handlerSt{
		serviceManager:  serviceManager,
		ApiAuthService:  apiAuthService,
		defaultResponse: dto.ErrorResponse{Message: "ok"},
	}
}

func (h *handlerSt) UpdatePrice(ctx fiber.Ctx) error {
	var inputBody dto.EditPriceRequest
	ctx.Bind().Body(&inputBody)
	userID := ctx.Locals("userID").(uint)

	h.serviceManager.SetConfigs(userID,
		inputBody.ReqHeaderEntry,
		inputBody.Prices,
		inputBody.Dates,
	)
	response := h.serviceManager.PriceUpdate()
	return ctx.JSON(response)
	// go h.serviceManager.PriceUpdate()
	// return ctx.JSON(h.defaultResponse)
}

func (h *handlerSt) UpdateDiscount(ctx fiber.Ctx) error {
	var inputBody dto.EditDiscountRequest
	ctx.Bind().Body(&inputBody)
	userID := ctx.Locals("userID").(uint)

	h.serviceManager.SetConfigs(userID,
		inputBody.ReqHeaderEntry,
		inputBody.Sites,
		inputBody.Dates,
	)
	response := h.serviceManager.DiscountUpdate(inputBody.DiscountPercent)
	return ctx.JSON(response)
	// go h.serviceManager.DiscountUpdate(inputBody.DiscountPercent)
	// return ctx.JSON(h.defaultResponse)
}

func (h *handlerSt) UpdateMinNight(ctx fiber.Ctx) error {
	var inputBody dto.EditMinNightRequest
	ctx.Bind().Body(&inputBody)
	userID := ctx.Locals("userID").(uint)

	h.serviceManager.SetConfigs(userID,
		inputBody.ReqHeaderEntry,
		inputBody.Sites,
		inputBody.Dates,
	)
	response := h.serviceManager.MinNightUpdate(inputBody.LimitDays)
	return ctx.JSON(response)
	// go h.serviceManager.MinNightUpdate(inputBody.LimitDays)
	// return ctx.JSON(h.defaultResponse)
}

func (h *handlerSt) UpdateCalendar(ctx fiber.Ctx) error {
	var inputBody dto.EditCalendarRequest
	ctx.Bind().Body(&inputBody)
	userID := ctx.Locals("userID").(uint)

	h.serviceManager.SetConfigs(userID,
		inputBody.ReqHeaderEntry,
		inputBody.Sites,
		inputBody.Dates,
	)
	response := h.serviceManager.CalendarUpdate(inputBody.Action)
	return ctx.JSON(response)
	// go h.serviceManager.CalendarUpdate(inputBody.Action)
	// return ctx.JSON(h.defaultResponse)
}

func (h *handlerSt) CheckAuth(ctx fiber.Ctx) error {
	var inputBody dto.ReqHeaderEntry
	errResponse, err := pkg.ValidateRequestBody(&inputBody, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errResponse)
	}

	userID := ctx.Locals("userID").(uint)
	h.serviceManager.SetConfigs(userID,
		inputBody,
		[]dto.SiteEntry{},
		[]string{},
	)
	response := h.serviceManager.CheckAuth()
	return ctx.JSON(response)
	// go h.serviceManager.CheckAuth()
	// return ctx.JSON(h.defaultResponse)
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
