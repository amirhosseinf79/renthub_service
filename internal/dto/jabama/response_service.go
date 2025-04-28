package jabama_dto

import "github.com/amirhosseinf79/renthub_service/internal/domain/models"

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
