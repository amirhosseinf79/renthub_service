package cloner

import (
	"errors"
	"fmt"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/internal/services/requests"
)

func (h *homsaService) handleUpdateResult(log *models.Log, body any, endpoint dto.EndP, userID uint, clientID, roomID string) (err error) {
	model, err := h.apiAuthRepo.GetByUnique(userID, clientID, h.service)
	if err != nil {
		log.FinalResult = err.Error()
		return err
	}
	url, err := h.getFullURL(endpoint, roomID)
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
	ok, res := request.Ok()
	fmt.Println("Request Result:", ok, res)
	if ok && h.service != "mihmansho" {
		log.FinalResult = "success"
		log.IsSucceed = true
		return nil
	}
	response := h.generateUpdateErrResponse()
	if response != nil {
		err = request.ParseInterface(response)
		if err != nil {
			log.FinalResult = err.Error()
			return err
		}
		ok, result := response.GetResult()
		if result == "" {
			result = "service error"
		}
		log.FinalResult = result
		if !ok {
			err = errors.New(result)
			return err
		}
		return err
	}
	return err
}
