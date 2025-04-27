package dto

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
