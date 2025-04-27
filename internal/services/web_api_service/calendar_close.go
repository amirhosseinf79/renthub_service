package cloner

import (
	"errors"

	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/internal/services/requests"
)

func (h *homsaService) CloseCalendar(fields dto.UpdateFields) error {
	endpoint := h.getEndpoints().CloseCalendar
	model, err := h.apiAuthRepo.GetByUnique(fields.UserID, fields.ClientID, h.service)
	if err != nil {
		return err
	}

	url, err := h.getFullURL(endpoint, fields.RoomID)
	if err != nil {
		return err
	}
	request := requests.New(endpoint.Method, url, h.getHeader(), h.getExtraHeader(model))
	body := h.generateCalendarBody(fields.RoomID, false, fields.Dates)
	request.Start(body, endpoint.ContentType)
	if !request.Ok() {
		return dto.ErrInvalidRequest
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
