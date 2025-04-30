package jajiga

import (
	"errors"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/internal/services/requests"
)

func (h *service) VerifyOtp(fields dto.RequiredFields, otp string) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID)
	endpoint := h.getEndpoints().LoginSecondStep
	model, err := h.apiAuthService.GetByUnique(fields.UserID, fields.ClientID, h.service)
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
	tokenModel := response.GetToken()
	tokenfields := dto.ApiAuthRequest{
		ClientID:     fields.ClientID,
		Username:     model.Username,
		Password:     otp,
		Service:      h.service,
		AccessToken:  tokenModel.AccessToken,
		RefreshToken: tokenModel.RefreshToken,
		Ucode:        tokenModel.Ucode,
	}
	err = h.apiAuthService.UpdateOrCreate(fields.UserID, tokenfields)
	if err != nil {
		log.FinalResult = err.Error()
		return log, err
	}
	log.FinalResult = "success"
	log.IsSucceed = true
	return log, err
}
