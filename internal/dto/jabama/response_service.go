package jabama_dto

import (
	"strings"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *Response) GetResult() (ok bool, result string) {
	ok = h.Error == nil
	result = "success"
	if !ok {
		result = h.Error.Message
	}
	return ok, result
}

func (h *Response) GetToken() *models.ApiAuth {
	return &models.ApiAuth{
		AccessToken:  h.Result.AccessToken,
		RefreshToken: h.Result.RefreshToken,
	}
}

func (h *UpdateErrorResponse) GetResult() (ok bool, result string) {
	if h.Success {
		return true, "success"
	}
	err := h.Error.Message
	if strings.Contains(h.Error.Message, "core-api") {
		err = dto.ErrTimeOut.Error()
	}
	return false, err
}

func (h *UpdateErrorResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
}
