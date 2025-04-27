package cloner

import (
	"errors"

	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/internal/services/requests"
)

func (h *homsaService) EditPricePerDays(fields dto.UpdateFields) error {
	endpoint := h.getEndpoints().EditPricePerDay
	model, err := h.apiAuthRepo.GetByUnique(fields.UserID, fields.ClientID, h.service)
	if err != nil {
		return err
	}

	url, err := h.getFullURL(endpoint, fields.RoomID)
	if err != nil {
		return err
	}
	request := requests.New(endpoint.Method, url, h.getHeader(), h.getExtraHeader(model))
	body := h.generatePriceBody(fields.RoomID, fields.Amount, fields.Dates)
	err = request.Start(body, endpoint.ContentType)
	if err != nil {
		return err
	}
	ok, result := request.Ok()
	if !ok {
		return result
	}
	response := h.generateMihmanshoErrResponse()
	if response != nil {
		request.ParseInterface(response)
		ok, result := response.GetResult()
		if !ok {
			return errors.New(result)
		}
	}
	return nil
}
