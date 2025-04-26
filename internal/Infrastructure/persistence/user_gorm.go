package persistence

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/domain/repository"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *userRepository) GetByID(id uint) (user *models.User, err error) {
	err = r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return
}

func (r *userRepository) GetByEmail(email string) (user *models.User, err error) {
	err = r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return
}

func (r *userRepository) GetAllByFilter(filter dto.UserFilter) (users []*models.User, total int64, err error) {
	query := r.db.Model(&models.User{})
	if filter.Email != "" {
		query = query.Where("email = ?", filter.Email)
	}
	err = query.Count(&total).Find(&users).Error
	return
}

func (r *userRepository) CheckEmailExists(email string) (exists bool, err error) {
	var count int64
	err = r.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
