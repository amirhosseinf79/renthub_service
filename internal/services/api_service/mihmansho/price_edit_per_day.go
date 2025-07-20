package mihmansho

import (
	"strconv"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	mihmansho_dto "github.com/amirhosseinf79/renthub_service/internal/dto/mihmansho"
)

func (h *service) EditPricePerDays(fields dto.UpdateFields) (log *models.Log, err error) {
	calendarResponse := mihmansho_dto.CalendarDetailsResponse{}
	log, err = h.GetCalendarDetails(fields, calendarResponse)
	if err != nil {
		return
	}

	guestPrice := -1
	for _, data := range calendarResponse.CalendarData {
		addedDay, _ := strconv.Atoi(data.AddedPrice)
		if addedDay > guestPrice {
			guestPrice = addedDay
		}
	}

	if guestPrice < 0 {
		err = dto.ErrInvalidRequest
		return
	}

	endpoint := h.getEndpoints().EditPricePerDay
	body := h.generatePriceBody(fields.Dates)
	err = h.handleUpdateResult(log, body, endpoint, fields, guestPrice)
	return log, err
}
