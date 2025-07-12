package interfaces

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	request_v1 "github.com/amirhosseinf79/renthub_service/internal/dto/request/v1"
	request_v2 "github.com/amirhosseinf79/renthub_service/internal/dto/request/v2"
)

type WebhookService interface {
	SendResult(response request_v1.ClientUpdateBody) (*models.Log, error)
	RefreshToken(userID uint) (*models.Log, error)
}

type WebhookService_v2 interface {
	SendResult(response request_v2.ClientUpdateBody) (*models.Log, error)
	RefreshToken(userID uint) (*models.Log, error)
}
