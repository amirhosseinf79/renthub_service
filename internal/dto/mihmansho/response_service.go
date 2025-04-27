package mihmansho_dto

import "github.com/amirhosseinf79/renthub_service/internal/domain/models"

func (m *MihmanshoProfileResponse) GetResult() (bool, string) {
	if m.UserInfo.Id == 0 {
		return false, "user not found"
	}
	return true, "success"
}

func (m *MihmanshoProfileResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
}

func (m *MihmanshoErrorResponse) GetResult() (bool, string) {
	if m.ErrorCode != 0 {
		return false, m.ErrorDescription
	}
	return true, "success"
}

func (m *MihmanshoErrorResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
}
