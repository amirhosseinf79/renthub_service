package mihmansho

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	mihmansho_dto "github.com/amirhosseinf79/renthub_service/internal/dto/mihmansho"
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
	roomDetails := mihmansho_dto.RoomDetailsResponse{}
	for counter < 2 {
		counter++
		log, err = h.GetRoomDetails(fields, &roomDetails)
	}
	if err != nil {
		log.FinalResult = err.Error()
		return
	}
	if roomDetails.Details.PaymentNote != "" {
		text := roomDetails.Details.PaymentNote
		cleanText := strings.ReplaceAll(text, ",", "")
		re := regexp.MustCompile(`\d+`)
		matches := re.FindAllString(cleanText, -1)
		for _, match := range matches {
			realNum, _ := strconv.Atoi(match)
			if realNum > 1000 {
				guestPrice = realNum / 1000
				break
			}
		}
	} else {
		guestPrice = 0
	}
	return
}
