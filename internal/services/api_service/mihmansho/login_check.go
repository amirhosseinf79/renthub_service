package mihmansho

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *service) CheckLogin(fields dto.RequiredFields) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID, dto.LoginCheck)
	endpoint := h.getEndpoints().GetProfile
	err = h.handleUpdateResult(log, nil, endpoint, dto.UpdateFields{RequiredFields: fields})
	return log, err
}
