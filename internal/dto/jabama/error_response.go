package jabama_dto

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/services/error_manager"
)

type msg struct {
	Message string `json:"message"`
}

type UpdateErrorResponse struct {
	Result  *[]string `json:"result"`
	Error   msg       `json:"error"`
	Success bool      `json:"success"`
}

func (h *UpdateErrorResponse) GetResult() (ok bool, result string) {
	if h.Success {
		return true, "success"
	}
	err := error_manager.ErrorLocalization(h.Error.Message)
	return false, err
}

func (h *UpdateErrorResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
}
