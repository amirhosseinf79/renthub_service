package repository

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

type UserRepository interface {
	GetByID(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetAllByFilter(filter dto.UserFilter) ([]*models.User, int64, error)
	CheckEmailExists(email string) (bool, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id uint) error
}
