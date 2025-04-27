package cloner

import (
	"errors"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/internal/services/requests"
)

func (h *homsaService) RemoveDiscount(fields dto.UpdateFields) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID)
	endpoint := h.getEndpoints().RemoveDiscount
	model, err := h.apiAuthRepo.GetByUnique(fields.UserID, fields.ClientID, h.service)
	if err != nil {
		log.FinalResult = err.Error()
		return log, err
	}
	url, err := h.getFullURL(endpoint, fields.RoomID)
	if err != nil {
		log.FinalResult = err.Error()
		return log, err
	}
	request := requests.New(endpoint.Method, url, h.getHeader(), h.getExtraHeader(model), log)
	body := h.generateRemoveDiscountBody(fields.RoomID, fields.Dates)
	err = request.Start(body, endpoint.ContentType)
	if err != nil {
		log.FinalResult = err.Error()
		return log, err
	}
	ok, err := request.Ok()
	if !ok {
		log.FinalResult = err.Error()
		return log, err
	}
	response := h.generateMihmanshoErrResponse()
	if response != nil {
		err = request.ParseInterface(response)
		if err != nil {
			log.FinalResult = err.Error()
			return log, err
		}
		ok, result := response.GetResult()
		if !ok {
			log.FinalResult = result
			return log, errors.New(result)
		}
	}
	log.FinalResult = "success"
	log.IsSucceed = true
	return log, err
}
