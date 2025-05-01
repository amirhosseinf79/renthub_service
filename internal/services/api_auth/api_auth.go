package apiauth

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/domain/repository"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

type apiAuthService struct {
	apiAuthRepo repository.ApiAuthRepository
}

func NewApiAuthService(apiAuthRepo repository.ApiAuthRepository) interfaces.ApiAuthInterface {
	return &apiAuthService{apiAuthRepo: apiAuthRepo}
}

func (h *apiAuthService) GetByUnique(userID uint, clientID string, service string) (*models.ApiAuth, error) {
	return h.apiAuthRepo.GetByUnique(userID, clientID, service)
}

func (h *apiAuthService) UpdateOrCreate(userID uint, fields dto.ApiAuthRequest) error {
	var err error
	exists := h.apiAuthRepo.CheckExists(userID, fields.ClientID, fields.Service)
	if exists {
		apiM, err := h.apiAuthRepo.GetByUnique(userID, fields.ClientID, fields.Service)
		if err != nil {
			return err
		}
		apiM.Username = fields.Username
		apiM.Password = fields.Password
		apiM.AccessToken = fields.AccessToken
		apiM.RefreshToken = fields.RefreshToken
		apiM.Ucode = fields.Ucode
		err = h.apiAuthRepo.Update(apiM)
		if err != nil {
			return err
		}
	} else {
		model := &models.ApiAuth{
			UserID:       userID,
			ClientID:     fields.ClientID,
			Service:      fields.Service,
			Username:     fields.Username,
			Password:     fields.Password,
			AccessToken:  fields.AccessToken,
			RefreshToken: fields.RefreshToken,
			Ucode:        fields.Ucode,
		}
		err = h.apiAuthRepo.Create(model)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *apiAuthService) GetClientAll(userID uint, clientID string) (list []*models.ApiAuth) {
	list = h.apiAuthRepo.GetAll(userID, clientID)
	return
}
