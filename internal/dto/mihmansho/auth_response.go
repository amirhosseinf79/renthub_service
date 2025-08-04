package mihmansho_dto

import "github.com/amirhosseinf79/renthub_service/internal/domain/models"

type AuthResponse struct {
	UCode            string `json:"ucode"`
	Token            string `json:"Token"`
	UserType         int    `json:"UserType"`
	ErrorCode        int    `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
}

func (m *AuthResponse) GetResult() (bool, string) {
	if m.ErrorCode != 0 {
		return false, m.ErrorDescription
	}
	return true, m.ErrorDescription
}

func (m *AuthResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{
		Ucode:       m.UCode,
		AccessToken: m.Token,
	}
}
