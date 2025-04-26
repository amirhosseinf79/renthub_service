package repository

import "github.com/amirhosseinf79/renthub_service/internal/domain/models"

type TokenRepository interface {
	GetByToken(token string) (*models.Token, error)
	GetByRefreshToken(refreshToken string) (*models.Token, error)
	Create(token *models.Token) error
	Update(token *models.Token) error
	Delete(token string) error
}
