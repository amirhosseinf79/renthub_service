package logger

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/domain/repository"
)

type loggerSt struct {
	logRepo repository.LogRepository
}

func NewLogger(logRepo repository.LogRepository) interfaces.LoggerInterface {
	return &loggerSt{logRepo: logRepo}
}

func (h *loggerSt) RecordLog(log *models.Log) error {
	return h.logRepo.Create(log)
}
