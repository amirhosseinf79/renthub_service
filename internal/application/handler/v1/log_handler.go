package handler_v1

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/gofiber/fiber/v3"
)

type logHandler struct {
	logService interfaces.LoggerInterface
}

func NewLogHandler(logService interfaces.LoggerInterface) interfaces.LoggerHandler {
	return &logHandler{logService: logService}
}

func (h *logHandler) GetLogs(c fiber.Ctx) error {
	var filters dto.LogFilters
	c.Bind().Query(&filters)
	userID := c.Locals("userID").(uint)
	response, errResp := h.logService.GetLogByfilter(userID, &filters)
	if errResp != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errResp)
	}
	return c.JSON(response)
}
