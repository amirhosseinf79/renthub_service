package cloner

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/internal/services/requests"
)

func (h *homsaService) CheckLogin(fields dto.RequiredFields) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID)
	endpoint := h.getEndpoints().GetProfile
	model, err := h.apiAuthRepo.GetByUnique(fields.UserID, fields.ClientID, h.service)
	if err != nil {
		log.FinalResult = err.Error()
		return log, err
	}
	url, err := h.getFullURL(endpoint)
	if err != nil {
		log.FinalResult = err.Error()
		return log, err
	}
	request := requests.New(endpoint.Method, url, h.getHeader(), h.getExtraHeader(model), log)
	err = request.Start(struct{}{}, endpoint.ContentType)
	if err != nil {
		log.FinalResult = err.Error()
		return log, err
	}
	ok, result := request.Ok()
	if !ok {
		log.FinalResult = result.Error()
		return log, err
	}

	response := h.generateProfileResponse()
	if response != nil {
		err = request.ParseInterface(response)
		if err != nil {
			log.FinalResult = err.Error()
			return log, err
		}
		ok, result := response.GetResult()
		if !ok {
			log.FinalResult = result
			return log, err
		}
	}
	log.FinalResult = "success"
	log.IsSucceed = true
	return log, err
}
