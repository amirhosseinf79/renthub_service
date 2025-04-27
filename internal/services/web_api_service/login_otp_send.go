package cloner

import (
	"errors"
	"fmt"

	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/internal/services/requests"
)

func (h *homsaService) SendOtp(fields dto.RequiredFields, phoneNumber string) error {
	if phoneNumber == "" {
		return dto.ErrEmptyPhone
	}

	url, err := h.getFullURL(h.getEndpoints().LoginFirstStep)
	if err != nil {
		return err
	}

	header := h.getHeader()
	body := h.generateSendOTPBody(phoneNumber)

	request := requests.New("POST", url, header, map[string]string{})
	err = request.BodyStart(body)
	if err != nil {
		return err
	}

	response := h.generateOTPResponse()
	if !request.Ok() {
		response = h.generateErrResponse()
	}
	err = request.ParseInterface(response)
	if err != nil {
		err = request.ParseInterface(response)
		if err != nil {
			return err
		}
	}
	fmt.Println(response.GetResult())
	ok, result := response.GetResult()
	if !ok {
		return errors.New(result)
	}

	record := dto.ApiEasyLogin{
		RequiredFields: fields,
		Username:       phoneNumber,
	}

	err = h.updateOrCreateAuthRecord(record, response.GetToken())
	if err != nil {
		return err
	}
	return nil
}
