package persistence

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/domain/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type tokenRepo struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) repository.TokenRepository {
	return &tokenRepo{db: db}
}

func (t *tokenRepo) GetByToken(token string) (*models.Token, error) {
	var tokenM models.Token
	// err := t.db.Where("token = ? AND updated_at >= ?", token, time.Now().Add(-8*time.Hour)).First(&tokenM).Error
	err := t.db.Where("token = ?", token).First(&tokenM).Error
	if err != nil {
		return nil, err
	}
	return &tokenM, nil
}

func (t *tokenRepo) GetByRefreshToken(refreshToken string) (*models.Token, error) {
	var tokenM models.Token
	err := t.db.Where("refresh_token = ?", refreshToken).First(&tokenM).Error
	if err != nil {
		return nil, err
	}
	return &tokenM, nil
}

func (t *tokenRepo) Create(token *models.Token) error {
	err := t.db.Omit(clause.Associations).Create(token).Error
	return err
}

func (t *tokenRepo) Update(token *models.Token) error {
	err := t.db.Omit(clause.Associations).Save(token).Error
	return err
}

func (t *tokenRepo) Delete(token string) error {
	err := t.db.Where("token = ?", token).Delete(&models.Token{}).Error
	return err
}
