package cloner

import (
	"errors"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/internal/services/requests"
)

func (h *homsaService) performLoginRequest(fields dto.ApiEasyLogin, otp bool) (authResponse interfaces.ApiResponseManager, err error) {
	endpoint := h.getEndpoints().LoginWithPass
	url, err := h.getFullURL(endpoint)
	if err != nil {
		return
	}
	header := h.getHeader()
	bodyRow := h.generateEasyLoginBody(fields)
	request := requests.New(endpoint.Method, url, header, map[string]string{})
	err = request.Start(bodyRow, endpoint.ContentType)
	if err != nil {
		return nil, err
	}
	response := h.generateAuthResponse()
	ok, _ := request.Ok()
	if !ok {
		response = h.generateErrResponse()
	}
	err = request.ParseInterface(response)
	if err != nil {
		return nil, err
	}
	ok, result := response.GetResult()
	if !ok {
		return nil, errors.New(result)
	}
	return response, nil
}

func (h *homsaService) updateOrCreateAuthRecord(fields dto.ApiEasyLogin, model *models.ApiAuth) error {
	var err error
	exists := h.apiAuthRepo.CheckExists(fields.UserID, fields.ClientID, h.service)
	if exists {
		apiM, err := h.apiAuthRepo.GetByUnique(fields.UserID, fields.ClientID, h.service)
		if err != nil {
			return err
		}
		apiM.Username = fields.Username
		apiM.Password = fields.Password
		apiM.AccessToken = model.AccessToken
		apiM.RefreshToken = model.RefreshToken
		apiM.Ucode = model.Ucode
		err = h.apiAuthRepo.Update(apiM)
		if err != nil {
			return err
		}
	} else {
		model := &models.ApiAuth{
			UserID:       fields.UserID,
			ClientID:     fields.ClientID,
			Service:      h.service,
			Username:     fields.Username,
			Password:     fields.Password,
			AccessToken:  model.AccessToken,
			RefreshToken: model.RefreshToken,
			Ucode:        model.Ucode,
		}
		err = h.apiAuthRepo.Create(model)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *homsaService) EasyLogin(fields dto.ApiEasyLogin) (err error) {
	response, err := h.performLoginRequest(fields, false)
	if err != nil {
		return
	}
	err = h.updateOrCreateAuthRecord(fields, response.GetToken())
	if err != nil {
		return
	}
	return nil
}
