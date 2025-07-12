package middleware_v2

import (
	"fmt"
	"regexp"
	"time"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	request_v2 "github.com/amirhosseinf79/renthub_service/internal/dto/request/v2"
	"github.com/amirhosseinf79/renthub_service/pkg"
	"github.com/gofiber/fiber/v3"
)

type validator struct{}

func NewValidator() interfaces.ValidatorInterface {
	return &validator{}
}

func (v *validator) DateCheck(c fiber.Ctx) error {
	var inputBody request_v2.DateEntry
	response := dto.ErrorResponse{
		Message: dto.ErrInvalidDate.Error(),
	}

	err := c.Bind().Body(&inputBody)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	fmt.Println("Dates:", inputBody.Dates)
	for _, date := range inputBody.Dates {
		_, err := time.Parse("2006-01-02", date)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}
		// oldDate := date.Before(time.Now())
		// if oldDate {
		// 	return c.Status(fiber.StatusBadRequest).JSON(response)
		// }
	}
	return c.Next()
}

func (v *validator) SendOTPCheck(c fiber.Ctx) error {
	var inputBody request_v2.OTPSendRequest
	response, err := pkg.ValidateRequestBody(&inputBody, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	regex, err := regexp.Compile("^09[0-9]{9}$")
	if err != nil {
		response := dto.ErrorResponse{
			Message: "Invalid regex pattern",
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	if !regex.Match([]byte(inputBody.PhoneNumebr)) {
		response := dto.ErrorResponse{
			Message: dto.ErrInvalidPhoneNumber.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	return c.Next()
}

func (v *validator) VerifyOTPCheck(c fiber.Ctx) error {
	var inputBody request_v2.OTPSendRequest
	response, err := pkg.ValidateRequestBody(&inputBody, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	regex, err := regexp.Compile("^09[0-9]{9}$")
	if err != nil {
		response := dto.ErrorResponse{
			Message: "Invalid regex pattern",
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	if !regex.Match([]byte(inputBody.PhoneNumebr)) {
		response := dto.ErrorResponse{
			Message: dto.ErrInvalidPhoneNumber.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	return c.Next()
}

func (v *validator) PriceUpdate(c fiber.Ctx) error {
	var inputBody request_v2.EditPriceRequest
	response, err := pkg.ValidateRequestBody(&inputBody, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	for _, site := range inputBody.Prices {
		if site.Price <= 0 {
			response := dto.ErrorResponse{
				Message: dto.ErrInvalidPrice.Error(),
			}
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}
	}
	return c.Next()
}

func (v *validator) RefReshTokenCheck(c fiber.Ctx) error {
	var inputBody request_v2.RefreshTokenRequest
	response, err := pkg.ValidateRequestBody(&inputBody, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	return c.Next()
}

func (v *validator) DiscountUpdate(c fiber.Ctx) error {
	var inputBody request_v2.EditDiscountRequest
	response, err := pkg.ValidateRequestBody(&inputBody, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	return c.Next()
}

func (v *validator) MinNightUpdate(c fiber.Ctx) error {
	var inputBody request_v2.EditMinNightRequest
	response, err := pkg.ValidateRequestBody(&inputBody, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	return c.Next()
}

func (v *validator) CalendarUpdate(c fiber.Ctx) error {
	var inputBody request_v2.EditCalendarRequest
	response, err := pkg.ValidateRequestBody(&inputBody, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	return c.Next()
}
