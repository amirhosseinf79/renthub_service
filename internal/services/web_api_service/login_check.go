package cloner

import (
	"errors"

	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/internal/services/requests"
)

func (h *homsaService) CheckLogin(fields dto.RequiredFields) error {
	endpoint := h.getEndpoints().GetProfile
	model, err := h.apiAuthRepo.GetByUnique(fields.UserID, fields.ClientID, h.service)
	if err != nil {
		return err
	}
	url, err := h.getFullURL(endpoint)
	if err != nil {
		return err
	}

	request := requests.New(endpoint.Method, url, h.getHeader(), h.getExtraHeader(model))
	err = request.Start(struct{}{}, endpoint.ContentType)
	if err != nil {
		return err
	}
	ok, result := request.Ok()
	if !ok {
		return result
	}

	response := h.generateProfileResponse()
	if response != nil {
		err := request.ParseInterface(response)
		if err != nil {
			return err
		}
		ok, result := response.GetResult()
		if !ok {
			return errors.New(result)
		}
	}
	return nil
}
