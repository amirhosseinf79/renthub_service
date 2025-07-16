package mihmansho

import (
	"errors"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"gorm.io/gorm"
)

func (h *service) AutoLogin(fields dto.RequiredFields) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID)
	model, err := h.apiAuthService.GetByUnique(fields.UserID, fields.ClientID, h.service)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = dto.ErrorApiTokenExpired
		}
		return log, err
	}
	// newSessionID, err := h.getSession(model.AccessToken, log)
	newSessionID, err := h.chromium.GetMihmanshoSessionID(model.AccessToken, log)
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
