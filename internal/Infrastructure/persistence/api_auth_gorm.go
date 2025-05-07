package persistence

import (
	"fmt"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/domain/repository"
	"gorm.io/gorm"
)

type apiAuthRepo struct {
	db *gorm.DB
}

func NewApiAuthRepository(db *gorm.DB) repository.ApiAuthRepository {
	return &apiAuthRepo{
		db: db,
	}
}

func (r *apiAuthRepo) CheckExists(userID uint, clientID string, service string) bool {
	var count int64
	r.db.Model(&models.ApiAuth{}).Where("user_id = ? AND client_id = ? AND service = ?", userID, clientID, service).Count(&count)
	return count > 0
}

func (r *apiAuthRepo) GetByUnique(userID uint, clientID string, service string) (model *models.ApiAuth, err error) {
	err = r.db.Where("user_id = ? AND client_id = ? AND service = ?", userID, clientID, service).First(&model).Error
	return
}

func (r *apiAuthRepo) GetAll(userID uint, clientID string) (list []*models.ApiAuth) {
	err := r.db.Where("user_id = ? AND client_id = ?", userID, clientID).Find(&list).Error
	fmt.Println("user_id = ? AND client_id = ?", userID, clientID, err)
	return
}

func (r *apiAuthRepo) Create(token *models.ApiAuth) error {
	return r.db.Create(token).Error
}

func (r *apiAuthRepo) Update(token *models.ApiAuth) error {
	return r.db.Save(token).Error
}
