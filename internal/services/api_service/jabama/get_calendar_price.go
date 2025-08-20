package jabama

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *service) GetCalendarPriceDetails(fields dto.UpdateFields, response any) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID, dto.GetCalendar)
	endpoint := h.getEndpoints().GetCalendarPrice
	body := h.generateCalendarDetailsBody(fields.Dates)
	err = h.handleUpdateResult(log, body, endpoint, fields, response)
	return
}
