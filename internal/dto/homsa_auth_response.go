package dto

import "github.com/amirhosseinf79/renthub_service/internal/domain/models"

type HomsaOTPResponse struct {
	Data homsaOTPData `json:"data"`
}

func (h *HomsaOTPResponse) GetResult() string {
	return "success"
}

func (h *HomsaOTPResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
}

type HomsaAuthResponse struct {
	UserID       int    `json:"user_id"`
	Email        string `json:"email"`
	FullName     string `json:"full_name"`
	ImageURL     string `json:"image_url"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpireAt     int    `json:"expire_at"`
	PhoneNumber  string `json:"phone_number"`
}

func (h *HomsaAuthResponse) GetResult() string {
	return "success"
}

func (h *HomsaAuthResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{
		AccessToken:  h.AccessToken,
		RefreshToken: h.RefreshToken,
	}
}

type HomsaErrorResponse struct {
	Code    string              `json:"code"`
	Message string              `json:"message"`
	Errors  map[string][]string `json:"errors"`
}

func (h *HomsaErrorResponse) GetResult() string {
	return h.Code
}

func (h *HomsaErrorResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
}
