package jajiga_dto

import "github.com/amirhosseinf79/renthub_service/internal/domain/models"

type OTPResponse struct {
	NotifyChannels []string `json:"notify_channels"`
}

func (r *OTPResponse) GetResult() (bool, string) {
	return true, "success"
}

func (r *OTPResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
}
