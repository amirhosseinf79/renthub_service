package homsa_dto

import "github.com/amirhosseinf79/renthub_service/internal/domain/models"

type HomsaErrorResponse struct {
	Code    string              `json:"code"`
	Message string              `json:"message"`
	Errors  map[string][]string `json:"errors"`
}

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
