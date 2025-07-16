package shab_dto

import (
	"fmt"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (r *AuthResponse) GetResult() (bool, string) {
	if r.Meta.Status >= 300 {
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
	isOk := r.Meta.Status < 300
	switch msgs := r.Meta.Messages.(type) {
	case map[string]any:
		for _, v := range msgs {
			if arr, ok := v.([]any); ok && len(arr) > 0 {
				if msg, ok := arr[0].(string); ok {
					return false, msg
				}
			}
		}
	case []any: // اگر پیام‌ها به صورت لیست ساده باشه
		for _, v := range msgs {
			if msg, ok := v.(string); ok {
				return false, msg
			}
		}
	case string:
		return false, msgs
	}
	return isOk, fmt.Sprintf("Error %v", r.Meta.Status)
}

func (r *ErrResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
}
