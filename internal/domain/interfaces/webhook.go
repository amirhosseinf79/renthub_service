package interfaces

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

type WebhookService interface {
	SendResult(response dto.ClientUpdateBody) (*models.Log, error)
	RefreshToken(userID uint) (*models.Log, error)
}
