package interfaces

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	request_v1 "github.com/amirhosseinf79/renthub_service/internal/dto/request/v1"
)

type WebhookService interface {
	SendResult(response request_v1.ClientUpdateBody) (*models.Log, error)
	RefreshToken(userID uint) (*models.Log, error)
}

type WebhookService_v2 interface {
	SendResult(response dto.WebhookFields) (*models.Log, error)
	RefreshToken(userID uint) (*models.Log, error)
}
