package interfaces

import "github.com/gofiber/fiber/v3"

type ManagerHandlerInterface interface {
	RefreshToken(ctx fiber.Ctx) error
	UpdatePrice(ctx fiber.Ctx) error
	UpdateDiscount(ctx fiber.Ctx) error
	UpdateMinNight(ctx fiber.Ctx) error
	UpdateCalendar(ctx fiber.Ctx) error
	TokenLogin(ctx fiber.Ctx) error
	CheckAuth(ctx fiber.Ctx) error
	SendServiceOTP(ctx fiber.Ctx) error
	VerifyServiceOTP(ctx fiber.Ctx) error
	SignOutClient(ctx fiber.Ctx) error
}
