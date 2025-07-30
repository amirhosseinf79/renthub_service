package auth_dto

import "github.com/amirhosseinf79/renthub_service/internal/domain/models"

type authResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	Ucode        string `json:"uCode,omitempty"`
}

func NewResponse(d *models.ApiAuth) authResponse {
	return authResponse{
		Token:        d.AccessToken,
		RefreshToken: d.RefreshToken,
		Ucode:        d.Ucode,
	}
}
