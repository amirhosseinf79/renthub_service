package mihmansho

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *service) AutoLogin(fields dto.RequiredFields) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID)
	model, err := h.apiAuthService.GetByUnique(fields.UserID, fields.ClientID, h.service)
	if err != nil {
		log.FinalResult = err.Error()
		return
	}
	newSessionID, err := h.getSession(model.AccessToken)
	if err != nil {
		log.FinalResult = err.Error()
		return
	}
	tokenfields := dto.ApiAuthRequest{
		ClientID:     fields.ClientID,
		Username:     model.Username,
		Password:     model.Password,
		Service:      h.service,
		AccessToken:  model.AccessToken,
		RefreshToken: newSessionID,
		Ucode:        model.Ucode,
	}
	err = h.apiAuthService.UpdateOrCreate(fields.UserID, tokenfields)
	if err != nil {
		log.FinalResult = err.Error()
		return log, err
	}
	log.FinalResult = "success"
	log.IsSucceed = true
	return log, err
}
