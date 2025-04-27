package cloner

import (
	"errors"

	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/internal/services/requests"
)

func (h *homsaService) VerifyOtp(fields dto.RequiredFields, otp string) (err error) {
	endpoint := h.getEndpoints().LoginSecondStep
	model, err := h.apiAuthRepo.GetByUnique(fields.UserID, fields.ClientID, h.service)
	if err != nil {
		return
	}
	url, err := h.getFullURL(endpoint)
	if err != nil {
		return
	}
	header := h.getHeader()
	bodyRow := h.generateVerifyOTPBody(model.Username, otp)
	request := requests.New(endpoint.Method, url, header, map[string]string{})
	err = request.Start(bodyRow, endpoint.ContentType)
	if err != nil {
		return err
	}
	response := h.generateAuthResponse()
	ok, _ := request.Ok()
	if !ok {
		response = h.generateErrResponse()
	}
	err = request.ParseInterface(response)
	if err != nil {
		return err
	}
	ok, result := response.GetResult()
	if !ok {
		return errors.New(result)
	}
	field := dto.ApiEasyLogin{
		RequiredFields: fields,
		Username:       model.Username,
		Password:       otp,
	}
	err = h.updateOrCreateAuthRecord(field, response.GetToken())
	if err != nil {
		return
	}
	return nil
}
