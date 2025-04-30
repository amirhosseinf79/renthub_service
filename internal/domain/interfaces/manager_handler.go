package interfaces

import "github.com/gofiber/fiber/v3"

type ManagerHandlerInterface interface {
	UpdatePrice(ctx fiber.Ctx) error
	UpdateDiscount(ctx fiber.Ctx) error
	UpdateMinNight(ctx fiber.Ctx) error
	UpdateCalendar(ctx fiber.Ctx) error
	TokenLogin(ctx fiber.Ctx) error
}
