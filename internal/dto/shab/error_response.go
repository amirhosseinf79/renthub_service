package shab_dto

import (
	"fmt"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/services/error_manager"
)

type ErrResponse struct {
	Meta meta `json:"meta"`
}

func (r *ErrResponse) GetResult() (bool, string) {
	isOk := r.Meta.Status < 300
	switch msgs := r.Meta.Messages.(type) {
	case map[string]any:
		for _, v := range msgs {
			if arr, ok := v.([]any); ok && len(arr) > 0 {
				if msg, ok := arr[0].(string); ok {
					return false, error_manager.ErrorLocalization(msg)
				}
			}
		}
	case []any: // اگر پیام‌ها به صورت لیست ساده باشه
		for _, v := range msgs {
			if msg, ok := v.(string); ok {
				return false, error_manager.ErrorLocalization(msg)
			}
		}
	case string:
		return false, error_manager.ErrorLocalization(msgs)
	}
	return isOk, fmt.Sprintf("Error %v", r.Meta.Status)
}

func (r *ErrResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
}
