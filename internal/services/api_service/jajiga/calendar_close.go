package jajiga

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *service) CloseCalendar(fields dto.UpdateFields) (log *models.Log, err error) {
	endpoint := h.getEndpoints().CloseCalendar
	log = h.initLog(fields.UserID, fields.ClientID)
	body := h.generateCalendarBody(fields.RoomID, false, fields.Dates)
	err = h.handleUpdateResult(log, body, endpoint, fields)
	return log, err
}
