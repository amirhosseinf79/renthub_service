package interfaces

import "github.com/gofiber/fiber/v3"

type ValidatorInterface interface {
	DateCheck(c fiber.Ctx) error
	PriceUpdate(c fiber.Ctx) error
	CalendarUpdate(c fiber.Ctx) error
	DiscountUpdate(c fiber.Ctx) error
	MinNightUpdate(c fiber.Ctx) error
	RefReshTokenCheck(c fiber.Ctx) error
	SendOTPCheck(c fiber.Ctx) error
	VerifyOTPCheck(c fiber.Ctx) error
	PaginationValidator(c fiber.Ctx) error
}

type ValidatorInterface_v2 interface {
	DateCheck(c fiber.Ctx) error
	PriceUpdate(c fiber.Ctx) error
	CalendarUpdate(c fiber.Ctx) error
	DiscountUpdate(c fiber.Ctx) error
	MinNightUpdate(c fiber.Ctx) error
	RefReshTokenCheck(c fiber.Ctx) error
	SendOTPCheck(c fiber.Ctx) error
	VerifyOTPCheck(c fiber.Ctx) error
	SignOutValidator(c fiber.Ctx) error
	RecieveDataValidator(c fiber.Ctx) error
}
