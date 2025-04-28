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
	// for _, value := range r.Meta.Messages {
	// 	if len(value) > 0 {
	// 		return false, value[0]
	// 	}
	// }
	if r.Meta.Status >= 300 {
		return false, fmt.Sprintf("Error %v", r.Meta.Status)
	}
	return true, "success"
}

func (r *ErrResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
}
