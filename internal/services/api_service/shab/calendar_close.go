package shab

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *service) CloseCalendar(fields dto.UpdateFields) (log *models.Log, err error) {
	endpoint := h.getEndpoints().CloseCalendar
	log = h.initLog(fields.UserID, fields.ClientID, dto.CloseCalendar)
	body := h.generateCalendarBody(false, fields.Dates)
	err = h.handleUpdateResult(log, body, endpoint, fields)
	return log, err
}
