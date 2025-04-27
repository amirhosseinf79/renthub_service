package homsa

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

	header := h.getHeader()
	url := h.getFullURL(h.getEndpoints().LoginFirstStep)
	body := h.generateSendOTPBody(phoneNumber)

	request := requests.New("POST", url, header, map[string]string{})
	err := request.BodyStart(body)
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
	if response.GetResult() != "success" {
		return errors.New(response.GetResult())
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
