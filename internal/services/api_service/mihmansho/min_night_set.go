package mihmansho

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *service) SetMinNight(fields dto.UpdateFields) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID)
	endpoint := h.getEndpoints().SetMinNight
	body := h.generateMinNightBody(fields.RoomID, fields.Dates, fields.Amount)
	err = h.handleUpdateResult(log, body, endpoint, fields)
	return log, err
}
