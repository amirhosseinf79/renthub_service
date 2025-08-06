package jajiga_dto

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/services/error_manager"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func (r *ErrorResponse) GetResult() (bool, string) {
	return false, error_manager.ErrorLocalization(r.Message)
}

func (r *ErrorResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
}
