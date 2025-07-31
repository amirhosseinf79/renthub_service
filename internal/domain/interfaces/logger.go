package interfaces

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	response_dto "github.com/amirhosseinf79/renthub_service/internal/dto/response"
	"github.com/gofiber/fiber/v3"
)

type LoggerInterface interface {
	RecordLog(log *models.Log) error
	GetLogByfilter(userID uint, filter *dto.LogFilters) (*response_dto.ListResponse[*models.Log], *dto.ErrorResponse)
}

type LoggerHandler interface {
	GetLogs(fiber.Ctx) error
}
