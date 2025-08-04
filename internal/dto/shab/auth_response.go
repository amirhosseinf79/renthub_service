package shab_dto

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

type authOk struct {
	AccessToken string `json:"access_token"`
}

type AuthResponse struct {
	Data authOk `json:"data"`
	Meta meta   `json:"meta"`
}

func (r *AuthResponse) GetResult() (bool, string) {
	if r.Meta.Status >= 300 {
		return false, dto.ErrInvalidRequest.Error()
	}
	return true, "success"
}

func (r *AuthResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{
		AccessToken: r.Data.AccessToken,
	}
}
