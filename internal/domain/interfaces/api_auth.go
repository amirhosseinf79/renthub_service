package interfaces

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

type ApiAuthInterface interface {
	GetByUnique(userID uint, clientID string, service string) (*models.ApiAuth, error)
	GetClientAll(userID uint, clientID string) []*models.ApiAuth
	UpdateOrCreate(userID uint, fields dto.ApiAuthRequest) error
}
