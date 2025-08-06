package otaghak_dto

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/services/error_manager"
)

type ErrorResponse struct {
	Message          string `json:"message"`
	Code             string `json:"code"`
	TechnicalMessage string `json:"technical_message"`
	HttpResponseCode int    `json:"http_response_code"`
	CorrelationId    string `json:"correlation_id"`
	Detail           string `json:"detail"`
}

func (r *ErrorResponse) GetResult() (bool, string) {
	return false, error_manager.ErrorLocalization(r.Message)
}

func (r *ErrorResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
}
