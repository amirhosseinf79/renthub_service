package otaghak_dto

import "github.com/amirhosseinf79/renthub_service/internal/domain/models"

type AuthOkResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func (r *AuthOkResponse) GetResult() (bool, string) {
	return true, "success"
}

func (r *AuthOkResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{
		AccessToken: r.AccessToken,
	}
}
