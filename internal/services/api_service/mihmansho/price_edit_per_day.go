package mihmansho

import (
	"strconv"
	"strings"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	mihmansho_dto "github.com/amirhosseinf79/renthub_service/internal/dto/mihmansho"
)

func (h *service) EditPricePerDays(fields dto.UpdateFields) (log *models.Log, err error) {
	calendarResponse := mihmansho_dto.CalendarDetailsResponse{}
	log, err = h.GetCalendarDetails(fields, &calendarResponse)
	if err != nil {
		log.FinalResult = err.Error()
		return
	}

	if len(fields.Dates) == 0 {
		err = dto.ErrInvalidRequest
		log.FinalResult = err.Error()
		return
	}

	firstDate := fields.Dates[0]
	guestPrice := -1
	for _, data := range calendarResponse.CalendarData {
		dateSection := strings.Split(firstDate, "/")
		day, _ := strconv.Atoi(dateSection[2])
		if data.Day == day {
			currPrice, _ := strconv.Atoi(data.AddedPrice)
			guestPrice = currPrice
			break
		}
	}

	if guestPrice < 0 {
		err = dto.ErrInvalidRequest
		log.FinalResult = err.Error()
		return
	}

	endpoint := h.getEndpoints().EditPricePerDay
	body := h.generatePriceBody(fields.Dates)
	err = h.handleUpdateResult(log, body, endpoint, fields, guestPrice)
	return log, err
}
