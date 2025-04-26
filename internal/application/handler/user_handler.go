package handler

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/pkg"
	"github.com/gofiber/fiber/v3"
)

type userHandler struct {
	userService interfaces.UserService
}

func NewUserHandler(userService interfaces.UserService) interfaces.UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) RegisterUser(c fiber.Ctx) error {
	var body dto.UserRegister
	response, err := pkg.ValidateRequestBody(&body, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	token, err := h.userService.RegisterUser(body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(token)
}

func (h *userHandler) LoginUser(c fiber.Ctx) error {
	var body dto.UserLogin
	response, err := pkg.ValidateRequestBody(&body, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	token, err := h.userService.LoginUser(body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(token)
}
