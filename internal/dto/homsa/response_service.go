package homsa_dto

import (
	"fmt"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
)

// Error Response
func (h *HomsaErrorResponse) GetResult() (bool, string) {
	for _, err := range h.Errors {
		if len(err) > 0 {
			return false, err[0]
		}
	}
	return false, h.Message
}

func (h *HomsaErrorResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
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

// OTP
func (h *HomsaOTPResponse) GetResult() (bool, string) {
	if h.Data.New {
		return true, "success"
	}
	return false, fmt.Sprintf("please wait %vs", h.Data.TTL)
}

func (h *HomsaOTPResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
}
