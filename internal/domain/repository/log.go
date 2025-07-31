package repository

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

type LogRepository interface {
	Create(*models.Log) error
	GetByFilter(userID uint, filter *dto.LogFilters) (int64, []*models.Log, error)
}
