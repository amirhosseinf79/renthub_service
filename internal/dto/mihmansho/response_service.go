package mihmansho_dto

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (m *MihmanshoProfileResponse) GetResult() (bool, string) {
	if m.UserInfo.Id == 0 {
		return false, dto.ErrUserNotFound.Error()
	}
	return true, m.UserInfo.UserName
}

func (m *MihmanshoProfileResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
}

func (m *MihmanshoErrorResponse) GetResult() (bool, string) {
	if m.ErrorCode != 0 {
		return false, m.ErrorDescription
	}
	return true, m.ErrorDescription
}

func (m *MihmanshoErrorResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
}

func (m *AuthResponse) GetResult() (bool, string) {
	if m.ErrorCode != 0 {
		return false, m.ErrorDescription
	}
	return true, m.ErrorDescription
}

func (m *AuthResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{
		Ucode:       m.UCode,
		AccessToken: m.Token,
	}
}
