package otaghak_dto

import "github.com/amirhosseinf79/renthub_service/internal/domain/models"

type OTPResponse struct {
	Odata           string `json:"@odata.context"`
	TimeExpiration  int    `json:"timeExpiration"`
	CellPhoneNumber string `json:"cellPhoneNumber"`
}

func (r *OTPResponse) GetResult() (bool, string) {
	return true, "success"
}

func (r *OTPResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
}
