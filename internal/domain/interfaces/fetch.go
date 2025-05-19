package interfaces

import "github.com/amirhosseinf79/renthub_service/internal/domain/models"

type FetchService interface {
	New(method, url string, headers, extra map[string]string, logger *models.Log) FetchService
	Start(body any, contentType string) error
	ParseInterface(any) error
	Ok() (bool, error)
}
