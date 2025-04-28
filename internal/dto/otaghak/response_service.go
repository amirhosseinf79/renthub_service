package otaghak_dto

import "github.com/amirhosseinf79/renthub_service/internal/domain/models"

func (r *AuthOkResponse) GetResult() (bool, string) {
	return true, "success"
}

func (r *AuthOkResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{
		AccessToken: r.AccessToken,
	}
}

func (r *OTPResponse) GetResult() (bool, string) {
	return true, "success"
}

func (r *OTPResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
}

func (r *ErrorResponse) GetResult() (bool, string) {
	return false, r.Message
}

func (r *ErrorResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
}
