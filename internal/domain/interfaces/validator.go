package interfaces

import "github.com/gofiber/fiber/v3"

type ValidatorInterface interface {
	DateCheck(c fiber.Ctx) error
	PriceUpdate(c fiber.Ctx) error
	CalendarUpdate(c fiber.Ctx) error
	DiscountUpdate(c fiber.Ctx) error
	MinNightUpdate(c fiber.Ctx) error
}
