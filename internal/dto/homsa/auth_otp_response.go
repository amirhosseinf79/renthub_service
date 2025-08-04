package homsa_dto

import (
	"fmt"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
)

type HomsaOTPResponse struct {
	Data homsaOTPData `json:"data"`
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
