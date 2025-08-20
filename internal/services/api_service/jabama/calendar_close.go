package jabama

import (
	"slices"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	jabama_dto "github.com/amirhosseinf79/renthub_service/internal/dto/jabama"
)

func (h *service) CloseCalendar(fields dto.UpdateFields) (log *models.Log, err error) {
	endpoint := h.getEndpoints().CloseCalendar
	log = h.initLog(fields.UserID, fields.ClientID, dto.CloseCalendar)
	body := h.generateCalendarBody(fields.Dates)
	finalResult := jabama_dto.UpdateCalendarStatusResponse{}
	err = h.handleUpdateResult(log, body, endpoint, fields, &finalResult)
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
			}
		}
		if err != nil {
			break
		}
	}
	return log, err
}
