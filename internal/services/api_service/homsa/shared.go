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
	if len(dates) > 1 {
		sort.Strings(dates)
	}
	return homsa_dto.HomsaCalendarBody{
		StartDate: dates[0],
		EndDate:   dates[len(dates)-1],
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
	if ok && h.service != "mihmansho" {
		log.FinalResult = "success"
		log.IsSucceed = true
		return nil
	} else if h.service != "mihmansho" {
		log.FinalResult = err.Error()
	}
	response := h.generateUpdateErrResponse()
	if response != nil {
		err2 := request.ParseInterface(response)
		if err2 == nil {
			ok, result := response.GetResult()
			if !ok && result != "" {
				err = errors.New(result)
				log.FinalResult = result
			} else if result != "" {
				log.FinalResult = result
			}
		}
	}
	return err
}
