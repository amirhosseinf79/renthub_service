package cloner

import (
	"errors"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	homsa_dto "github.com/amirhosseinf79/renthub_service/internal/dto/homsa"
	jabama_dto "github.com/amirhosseinf79/renthub_service/internal/dto/jabama"
	jajiga_dto "github.com/amirhosseinf79/renthub_service/internal/dto/jajiga"
	mihmansho_dto "github.com/amirhosseinf79/renthub_service/internal/dto/mihmansho"
	otaghak_dto "github.com/amirhosseinf79/renthub_service/internal/dto/otaghak"
	shab_dto "github.com/amirhosseinf79/renthub_service/internal/dto/shab"
	"github.com/amirhosseinf79/renthub_service/internal/services/requests"
)

func (h *homsaService) generateSendOTPBody(phoneNumber string) any {
	switch h.service {
	case "homsa":
		return homsa_dto.HomsaOTPLogin{
			Mobile: phoneNumber,
		}
	case "jabama":
		return jabama_dto.OTPLogin{
			Mobile: phoneNumber,
		}
	case "jajiga":
		return jajiga_dto.OTPLogin{
			Mobile: phoneNumber,
			ISO2:   "IR",
		}
	case "otaghak":
		return otaghak_dto.OTPBody{
			UserName:   phoneNumber,
			IsShortOtp: true,
		}
	case "shab":
		return shab_dto.OTPBody{
			Mobile:      phoneNumber,
			CountryCode: "+98",
		}
	case "mihmansho":
		return mihmansho_dto.OTPBody{
			Mobile: phoneNumber,
			IsCode: true,
		}
	}
	return nil
}

func (h *homsaService) SendOtp(fields dto.RequiredFields, phoneNumber string) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID)
	endpoint := h.getEndpoints().LoginFirstStep
	url, err := h.getFullURL(endpoint)
	if err != nil {
		log.FinalResult = err.Error()
		return log, err
	}
	header := h.getHeader()
	body := h.generateSendOTPBody(phoneNumber)
	request := requests.New(endpoint.Method, url, header, map[string]string{}, log)
	err = request.Start(body, endpoint.ContentType)
	if err != nil {
		log.FinalResult = err.Error()
		return log, err
	}
	response := h.generateOTPResponse()
	ok, _ := request.Ok()
	if !ok {
		response = h.generateErrResponse()
	}
	err = request.ParseInterface(response)
	if err != nil {
		log.FinalResult = err.Error()
		return log, err
	}
	ok, result := response.GetResult()
	log.FinalResult = result
	if !ok {
		return log, errors.New(result)
	}
	record := dto.ApiEasyLogin{
		RequiredFields: fields,
		Username:       phoneNumber,
	}
	err = h.updateOrCreateAuthRecord(record, response.GetToken())
	if err != nil {
		log.FinalResult = err.Error()
		return log, err
	}
	log.FinalResult = "success"
	log.IsSucceed = true
	return log, err
}
