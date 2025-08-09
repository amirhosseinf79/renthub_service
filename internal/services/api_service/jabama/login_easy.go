package jabama

import (
	"errors"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"gorm.io/gorm"
)

func (h *service) AutoLogin(fields dto.RequiredFields) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID, dto.AutoLogin)
	endpoint := h.getEndpoints().LoginWithPass
	model, err := h.apiAuthService.GetByUnique(fields.UserID, fields.ClientID, h.service)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = dto.ErrorApiTokenExpired
		}
		log.FinalResult = err.Error()
		return log, err
	}
	if model.RefreshToken == "" {
		err = dto.ErrorApiTokenExpired
		log.FinalResult = err.Error()
		return log, err
	}
	url, err := h.getFullURL(endpoint, model.RefreshToken)
	if err != nil {
		log.FinalResult = err.Error()
		return log, err
	}
	header := h.getHeader()
	request := h.request.New(endpoint.Method, url, header, map[string]string{}, log)
	err = request.Start(nil, endpoint.ContentType)
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
		Password:     model.Password,
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
