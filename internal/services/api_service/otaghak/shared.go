package otaghak

import (
	"errors"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"gorm.io/gorm"
)

func (h *service) handleUpdateResult(log *models.Log, body any, endpoint dto.EndP, fields dto.UpdateFields) (err error) {
	model, err := h.apiAuthService.GetByUnique(fields.UserID, fields.ClientID, h.service)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = dto.ErrorApiTokenExpired
		}
		log.FinalResult = err.Error()
		return err
	}
	url, err := h.getFullURL(endpoint, fields.RoomID, fields.Amount)
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

func (h *service) handleGetResult(log *models.Log, endpoint dto.EndP, fields dto.RecieveFields, response any) error {
	model, err := h.apiAuthService.GetByUnique(fields.UserID, fields.ClientID, h.service)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = dto.ErrorApiTokenExpired
		}
		log.FinalResult = err.Error()
		return err
	}

	url, err := h.getFullURL(endpoint)
	if err != nil {
		log.FinalResult = err.Error()
		return err
	}

	request := h.request.New(endpoint.Method, url, h.getHeader(), h.getExtraHeader(model), log)
	err = request.Start(fields.Filters, endpoint.ContentType)
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
	err = request.ParseInterface(response)
	if err != nil {
		log.FinalResult = err.Error()
		return err
	}
	log.FinalResult = "success"
	log.IsSucceed = true
	return nil
}
