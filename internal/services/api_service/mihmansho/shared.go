package mihmansho

import (
	"errors"
	"strconv"
	"strings"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	mihmansho_dto "github.com/amirhosseinf79/renthub_service/internal/dto/mihmansho"
	"github.com/amirhosseinf79/renthub_service/pkg"
	"gorm.io/gorm"
)

func (h *service) handleUpdateResult(log *models.Log, body any, endpoint dto.EndP, fields dto.UpdateFields, addedGuests int) (err error) {
	model, err := h.apiAuthService.GetByUnique(fields.UserID, fields.ClientID, h.service)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = dto.ErrorApiTokenExpired
		}
		log.FinalResult = err.Error()
		return err
	}
	url, err := h.getFullURL(endpoint, fields.RoomID, fields.Amount/1000, addedGuests)
	if err != nil {
		log.FinalResult = err.Error()
		return err
	}
	request := h.request.New(endpoint.Method, url, h.getHeader(), h.getExtraHeader(model), log)
	err = request.Start(body, endpoint.ContentType)
	if err != nil {
		log.FinalResult = err.Error()
		return err
	}
	response := h.generateUpdateErrResponse()
	err = request.ParseInterface(response)
	if err != nil {
		return err
	}
	ok, result := response.GetResult()
	log.FinalResult = result
	if !ok {
		err = errors.New(result)
	} else {
		log.IsSucceed = true
	}
	return err
}

func (h *service) getAddGuestPrice(fields dto.UpdateFields) (guestPrice int, log *models.Log, err error) {
	guestPrice = -1
	counter := 0
	found := false
	for !found && counter < 2 {
		counter++
		calendarResponse := mihmansho_dto.CalendarDetailsResponse{}
		log, err = h.GetCalendarDetails(fields, &calendarResponse)
		jdates := pkg.DatesToJalali(fields.Dates, false)
		if len(jdates) == 0 {
			return
		}

		firstDate := jdates[0]
		for _, data := range calendarResponse.CalendarData {
			dateSection := strings.Split(firstDate, "/")
			day, _ := strconv.Atoi(dateSection[2])
			if data.Day == day {
				currPrice, _ := strconv.Atoi(data.AddedPrice)
				guestPrice = currPrice
				found = true
				break
			}
		}
	}
	if err != nil {
		log.FinalResult = err.Error()
		return
	}
	if guestPrice < 0 {
		err = dto.ErrInvalidDate
		log.FinalResult = err.Error()
		return
	}
	return
}
