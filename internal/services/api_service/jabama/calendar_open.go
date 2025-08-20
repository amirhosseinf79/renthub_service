package jabama

import (
	"slices"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	jabama_dto "github.com/amirhosseinf79/renthub_service/internal/dto/jabama"
)

func (h *service) OpenCalendar(fields dto.UpdateFields) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID, dto.OpenCalendar)
	endpoint := h.getEndpoints().OpenCalendar
	body := h.generateCalendarBody(fields.Dates)
	finalResult := jabama_dto.UpdateCalendarStatusResponse{}
	err = h.handleUpdateResult(log, body, endpoint, fields, &finalResult)
	if err != nil {
		return log, err
	}
	for _, data := range finalResult.Result.Price.Custom {
		if slices.Contains(fields.Dates, data.Date.Format("2006-01-02")) {
			if data.Status != "available" {
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
