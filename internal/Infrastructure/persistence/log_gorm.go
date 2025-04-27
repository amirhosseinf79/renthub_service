package persistence

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/domain/repository"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"gorm.io/gorm"
)

type logRepo struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) repository.LogRepository {
	return &logRepo{
		db: db,
	}
}

func (r *logRepo) Create(model *models.Log) error {
	return r.db.Create(model).Error
}

func (r *logRepo) GetByFilter(filter *dto.RequiredFields, service *string) ([]*models.Log, error) {
	var logs []*models.Log
	model := r.db.Model(&models.Log{})
	if filter != nil {
		if filter.UserID > 0 {
			model = model.Where("user_id = ?", filter.UserID)
		}
		if filter.ClientID != "" {
			model = model.Where("client_id = ?", filter.ClientID)
		}
	}
	if service != nil {
		model.Where("service = ?", service)
	}
	err := model.Find(&logs).Error
	return logs, err
}
