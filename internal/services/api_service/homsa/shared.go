package homsa

import (
	"errors"
	"sort"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	homsa_dto "github.com/amirhosseinf79/renthub_service/internal/dto/homsa"
	"github.com/amirhosseinf79/renthub_service/internal/services/requests"
	"gorm.io/gorm"
)

func (h *service) generateCalendarBody(dates []string) any {
	cDates := make([]string, len(dates))
	copy(cDates, dates)
	if len(cDates) > 1 {
		sort.Strings(cDates)
	}
	return homsa_dto.HomsaCalendarBody{
		StartDate: cDates[0],
		EndDate:   cDates[len(cDates)-1],
	}
}

func (h *service) handleUpdateResult(log *models.Log, body any, endpoint dto.EndP, fields dto.UpdateFields) (err error) {
	model, err := h.apiAuthService.GetByUnique(fields.UserID, fields.ClientID, h.service)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = dto.ErrorUnauthorized
		}
		log.FinalResult = err.Error()
		return err
	}
	url, err := h.getFullURL(endpoint, fields.RoomID, fields.Amount)
	if err != nil {
		log.FinalResult = err.Error()
		return err
	}
	request := requests.New(endpoint.Method, url, h.getHeader(), h.getExtraHeader(model), log)
	err = request.Start(body, endpoint.ContentType)
	if err != nil {
		log.FinalResult = err.Error()
		return err
	}
	ok, err := request.Ok()
	if !ok {
		log.FinalResult = err.Error()
		response := h.generateUpdateErrResponse()
		err2 := request.ParseInterface(response)
		if err2 != nil {
			return err
		}
		_, result := response.GetResult()
		log.FinalResult = result
		err = errors.New(result)
		return err
	}
	log.FinalResult = "success"
	log.IsSucceed = true
	return nil
}
