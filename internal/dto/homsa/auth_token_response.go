package homsa_dto

import "github.com/amirhosseinf79/renthub_service/internal/domain/models"

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

// Auth Response
func (h *HomsaAuthResponse) GetResult() (bool, string) {
	return true, "success"
}

func (h *HomsaAuthResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{
		AccessToken:  h.AccessToken,
		RefreshToken: h.RefreshToken,
	}
}
