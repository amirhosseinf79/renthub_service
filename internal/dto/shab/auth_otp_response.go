package shab_dto

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

type AuthOTPResponse struct {
	Meta meta `json:"meta"`
}

func (r *AuthOTPResponse) GetResult() (bool, string) {
	if r.Meta.Status >= 300 {
		return false, dto.ErrInvalidRequest.Error()
	}
	return true, "success"
}

func (r *AuthOTPResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
}
