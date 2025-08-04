package otaghak_dto

import "github.com/amirhosseinf79/renthub_service/internal/domain/models"

type ErrorResponse struct {
	Message          string `json:"message"`
	Code             string `json:"code"`
	TechnicalMessage string `json:"technical_message"`
	HttpResponseCode int    `json:"http_response_code"`
	CorrelationId    string `json:"correlation_id"`
	Detail           string `json:"detail"`
}

func (r *ErrorResponse) GetResult() (bool, string) {
	return false, r.Message
}

func (r *ErrorResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
}
