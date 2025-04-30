package shab_dto

import (
	"fmt"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (r *AuthResponse) GetResult() (bool, string) {
	if r.Meta.Status > 300 {
		return false, dto.ErrInvalidRequest.Error()
	}
	return true, "success"
}

func (r *AuthResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{
		AccessToken: r.Data.AccessToken,
	}
}

func (r *AuthOTPResponse) GetResult() (bool, string) {
	if r.Meta.Status >= 300 {
		return false, dto.ErrInvalidRequest.Error()
	}
	return true, "success"
}

func (r *AuthOTPResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
}

func (r *ErrResponse) GetResult() (bool, string) {
	if msgs, ok := r.Meta.Messages.(map[string][]string); ok {
		for _, value := range msgs {
			if len(value) > 0 {
				return false, value[0]
			}
		}
	} else if msgs, ok := r.Meta.Messages.([]string); ok {
		for _, value := range msgs {
			return false, value
		}
	}
	return true, fmt.Sprintf("Error %v", r.Meta.Status)
}

func (r *ErrResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
}
