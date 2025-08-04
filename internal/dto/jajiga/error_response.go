package jajiga_dto

import "github.com/amirhosseinf79/renthub_service/internal/domain/models"

type ErrorResponse struct {
	Message string `json:"message"`
}

func (r *ErrorResponse) GetResult() (bool, string) {
	return false, r.Message
}

func (r *ErrorResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
}
