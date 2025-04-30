package interfaces

import "github.com/amirhosseinf79/renthub_service/internal/domain/models"

type LoggerInterface interface {
	RecordLog(log *models.Log) error
}
