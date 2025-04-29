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
	"github.com/google/uuid"
)

func (h *homsaService) generateVerifyOTPBody(phoneNumber string, code string) any {
	switch h.service {
	case "homsa":
		return homsa_dto.HomsaLoginUserPass{
			Mobile:   phoneNumber,
			Password: code,
			UseOTP:   true,
		}
	case "jabama":
		return jabama_dto.OTPLogin{
			Mobile: phoneNumber,
			Code:   code,
		}
	case "jajiga":
		return jajiga_dto.JajigaTokenAuthRequestBody{
			Mobile:   phoneNumber,
			Token:    &code,
			ClientID: uuid.New().String(),
			ISO2:     "IR",
		}
	case "otaghak":
		return otaghak_dto.OtaghakAuthRequestBody{
			UserName:     phoneNumber,
			ClientId:     "Otaghak",
			ClientSecret: "secret",
			ArcValues:    map[string]string{"OtpCode": code},
		}
	case "shab":
		return shab_dto.VerifyOTOBody{
			Mobile:      phoneNumber,
			CountryCode: "+98",
			Code:        code,
		}
	case "mihmansho":
		return mihmansho_dto.OTPVerifyBody{
			Mobile: phoneNumber,
			Code:   code,
		}
	default:
		return nil
	}
}

func (h *homsaService) VerifyOtp(fields dto.RequiredFields, otp string) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID)
	endpoint := h.getEndpoints().LoginSecondStep
	model, err := h.apiAuthRepo.GetByUnique(fields.UserID, fields.ClientID, h.service)
	if err != nil {
		log.FinalResult = err.Error()
		return log, err
	}
	url, err := h.getFullURL(endpoint)
	if err != nil {
		log.FinalResult = err.Error()
		return log, err
	}
	header := h.getHeader()
	bodyRow := h.generateVerifyOTPBody(model.Username, otp)
	request := requests.New(endpoint.Method, url, header, map[string]string{}, log)
	err = request.Start(bodyRow, endpoint.ContentType)
	if err != nil {
		log.FinalResult = err.Error()
		return log, err
	}
	response := h.generateAuthResponse()
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
	field := dto.ApiEasyLogin{
		RequiredFields: fields,
		Username:       model.Username,
		Password:       otp,
	}
	err = h.updateOrCreateAuthRecord(field, response.GetToken())
	if err != nil {
		log.FinalResult = err.Error()
		return log, err
	}
	log.FinalResult = "success"
	log.IsSucceed = true
	return log, err
}
