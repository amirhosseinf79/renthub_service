package persistence

import (
	"time"

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

func (r *logRepo) GetByFilter(userID uint, filter *dto.LogFilters) (int64, []*models.Log, error) {
	var logs []*models.Log
	model := r.db.Model(&models.Log{})
	model = model.Where("user_id = ?", userID)
	if filter != nil {
		if filter.Service != "" {
			model = model.Where("service = ?", filter.Service)
		}
		if filter.ClientID != "" {
			model = model.Where("client_id = ?", filter.ClientID)
		}
		if filter.UpdateID != "" {
			model = model.Where("update_id = ?", filter.UpdateID)
		}
		if filter.RoomID != "" {
			model = model.Where("request_body LIKE ?", "%"+filter.RoomID+"%")
		}
		if filter.Action != "" {
			model = model.Where("action = ?", filter.Action)
		}
		if filter.ResponseBody != "" {
			model = model.Where("response_body LIKE ?", "%"+filter.ResponseBody+"%")
		}
		if filter.IsSucceed != nil {
			model = model.Where("is_succeed = ?", filter.IsSucceed)
		}
		if filter.FromDate != "" {
			parsedTime, err := time.Parse("2006-01-02T15:04:05", filter.FromDate)
			if err != nil {
				return 0, nil, dto.ErrInvalidDate
			}
			model = model.Where("created_at >= ?", parsedTime.UTC())
		}
		if filter.ToDate != "" {
			parsedTime, err := time.Parse("2006-01-02T15:04:05", filter.ToDate)
			if err != nil {
				return 0, nil, dto.ErrInvalidDate
			}
			model = model.Where("created_at <= ?", parsedTime.UTC())
		}
	}
	var total int64
	model.Count(&total)
	if filter == nil || filter.PageSize == 0 || filter.PageSize > 100 {
		filter.PageSize = 10
	}
	model = model.Offset((int(filter.Page) - 1) * int(filter.PageSize)).Limit(int(filter.PageSize))
	err := model.Order("id DESC").Find(&logs).Error
	return total, logs, err
}
