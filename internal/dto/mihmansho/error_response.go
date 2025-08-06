package mihmansho_dto

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/internal/services/error_manager"
)

type MihmanshoErrorResponse struct {
	ResponseError    int    `json:"responseError"`
	ErrorCode        int    `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
}

func (m *MihmanshoErrorResponse) GetResult() (bool, string) {
	var msg string = "success"
	if m.ErrorCode != 0 || m.ResponseError != 0 {
		if m.ResponseError == 2 {
			msg = dto.ErrorSessionNotFound.Error()
		} else if m.ErrorDescription != "" {
			msg = m.ErrorDescription
		} else {
			msg = dto.ErrUnknownMsg.Error()
		}
		return false, error_manager.ErrorLocalization(msg)
	}
	return true, error_manager.ErrorLocalization(msg)
}

func (m *MihmanshoErrorResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
}
