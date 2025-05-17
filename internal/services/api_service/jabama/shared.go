package jabama

import (
	"errors"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/internal/services/requests"
	"gorm.io/gorm"
)

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

func (h *service) handleGet(log *models.Log, body any, endpoint dto.EndP, fields dto.GetDetail, response any) (err error) {
	model, err := h.apiAuthService.GetByUnique(fields.UserID, fields.ClientID, h.service)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = dto.ErrorUnauthorized
		}
		log.FinalResult = err.Error()
		return err
	}
	url, err := h.getFullURL(endpoint, fields.RoomID)
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
		if log.StatusCode == 401 {
			data := dto.RequiredFields{
				UserID:   fields.UserID,
				ClientID: fields.ClientID,
			}
			_, err2 := h.AutoLogin(data)
			if err2 != nil {
				log.FinalResult = err2.Error()
				return err2
			}
		}
		log.FinalResult = err.Error()
		return err
	}
	log.FinalResult = "success"
	err2 := request.ParseInterface(response)
	if err2 != nil {
		return err2
	}
	return nil
}
