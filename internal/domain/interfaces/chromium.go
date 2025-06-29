package interfaces

import "github.com/amirhosseinf79/renthub_service/internal/domain/models"

type ChromeService interface {
	Close()
	GetMihmanshoSessionID(token string, log *models.Log) (string, error)
	GetJajigaHeaders(log *models.Log) (map[string]string, error)
}
