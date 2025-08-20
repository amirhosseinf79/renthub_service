package jabama

import (
	"slices"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *service) CloseCalendar(fields dto.UpdateFields) (log *models.Log, err error) {
	endpoint := h.getEndpoints().CloseCalendar
	log = h.initLog(fields.UserID, fields.ClientID, dto.CloseCalendar)
	body := h.generateCalendarBody(fields.Dates)
	err = h.handleUpdateResult(log, body, endpoint, fields, nil)
	if err != nil {
		return log, err
	}
	log2, calendarResponse, err := h.GetCalendarDetails(fields)
	if err != nil {
		return log2, err
	}
	for _, data := range calendarResponse.Result.Calendar {
		if slices.Contains(fields.Dates, data.Date) {
			if data.Status != "disabledByHost" {
				err = dto.ErrUnknownMsg
				log.FinalResult = err.Error()
				log.IsSucceed = false
				return log2, err
			}
		}
	}
	return log, err
}
