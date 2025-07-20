package homsa

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *service) CloseCalendar(fields dto.UpdateFields) (log *models.Log, err error) {
	endpoint := h.getEndpoints().CloseCalendar
	log = h.initLog(fields.UserID, fields.ClientID, dto.CloseCalendar)
	bodies := h.generateCalendarBody(fields.Dates)
	for _, body := range bodies {
		err = h.handleUpdateResult(log, body, endpoint, fields)
		if err != nil {
			return log, err
		}
	}
	return log, err
}
