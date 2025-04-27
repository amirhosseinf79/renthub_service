package cloner

import (
	"errors"

	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/internal/services/requests"
)

func (h *homsaService) CheckLogin(fields dto.RequiredFields) error {
	model, err := h.apiAuthRepo.GetByUnique(fields.UserID, fields.ClientID, h.service)
	if err != nil {
		return err
	}
	url, err := h.getFullURL(h.getEndpoints().GetProfile)
	if err != nil {
		return err
	}

	request := requests.New("GET", url, h.getHeader(), h.getExtraHeader(model))
	err = request.BodyStart(struct{}{})
	if err != nil {
		return err
	}
	if !request.Ok() {
		return dto.ErrorUnauthorized
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
