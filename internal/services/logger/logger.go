package logger

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/domain/repository"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	response_dto "github.com/amirhosseinf79/renthub_service/internal/dto/response"
)

type loggerSt struct {
	logRepo repository.LogRepository
}

func NewLogger(logRepo repository.LogRepository) interfaces.LoggerInterface {
	return &loggerSt{logRepo: logRepo}
}

func (h *loggerSt) RecordLog(log *models.Log) error {
	if log != nil {
		return h.logRepo.Create(log)
	}
	return nil
}

func (h *loggerSt) GetLogByfilter(userID uint, filter *dto.LogFilters) (response *response_dto.ListResponse[*models.Log], errRespose *dto.ErrorResponse) {
	total, list, err := h.logRepo.GetByFilter(userID, filter)
	if err != nil {
		errRespose = &dto.ErrorResponse{
			Message: err.Error(),
		}
		return
	}

	response = response_dto.NewListResponse(total, filter.Page, filter.PageSize, list)
	return
}
