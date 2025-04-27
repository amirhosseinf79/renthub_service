package cloner

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/internal/services/requests"
)

func (h *homsaService) CloseCalendar(fields dto.UpdateFields) (log *models.Log) {
	log = h.initLog(fields.UserID, fields.ClientID)
	endpoint := h.getEndpoints().CloseCalendar
	model, err := h.apiAuthRepo.GetByUnique(fields.UserID, fields.ClientID, h.service)
	if err != nil {
		log.FinalResult = err.Error()
		return
	}
	url, err := h.getFullURL(endpoint, fields.RoomID)
	if err != nil {
		log.FinalResult = err.Error()
		return
	}
	request := requests.New(endpoint.Method, url, h.getHeader(), h.getExtraHeader(model), log)
	body := h.generateCalendarBody(fields.RoomID, false, fields.Dates)
	err = request.Start(body, endpoint.ContentType)
	if err != nil {
		log.FinalResult = err.Error()
		return
	}
	ok, result := request.Ok()
	if !ok {
		log.FinalResult = result.Error()
		return
	}
	response := h.generateMihmanshoErrResponse()
	if response != nil {
		err = request.ParseInterface(response)
		if err != nil {
			log.FinalResult = err.Error()
			return
		}
		ok, result := response.GetResult()
		if !ok {
			log.FinalResult = result
			return
		}
	}
	log.FinalResult = "success"
	log.IsSucceed = true
	return
}
