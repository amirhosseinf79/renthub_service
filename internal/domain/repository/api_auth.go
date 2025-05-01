package repository

import "github.com/amirhosseinf79/renthub_service/internal/domain/models"

type ApiAuthRepository interface {
	CheckExists(userID uint, clientID string, service string) bool
	GetByUnique(userID uint, clientID string, service string) (*models.ApiAuth, error)
	GetAll(userID uint, clientID string) []*models.ApiAuth
	Create(*models.ApiAuth) error
	Update(*models.ApiAuth) error
}
