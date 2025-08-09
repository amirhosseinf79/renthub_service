package interfaces

import "github.com/gofiber/fiber/v3"

type RecieveHandler interface {
	GetReservations(fiber.Ctx) error
}
