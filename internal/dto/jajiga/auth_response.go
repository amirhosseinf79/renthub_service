package jajiga_dto

import "github.com/amirhosseinf79/renthub_service/internal/domain/models"

type AuthOkResponse struct {
	JWTToken string `json:"jwt_token"`
}

func (r *AuthOkResponse) GetResult() (bool, string) {
	return true, "success"
}

func (r *AuthOkResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{
		AccessToken: r.JWTToken,
	}
}
