package jajiga

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *service) OpenCalendar(fields dto.UpdateFields) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID)
	endpoint := h.getEndpoints().OpenCalendar
	body := h.generateCalendarBody(fields.RoomID, true, fields.Dates)
	err = h.handleUpdateResult(log, body, endpoint, fields)
	return log, err
}
