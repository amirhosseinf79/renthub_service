package cloner

import (
	"errors"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	homsa_dto "github.com/amirhosseinf79/renthub_service/internal/dto/homsa"
	jajiga_dto "github.com/amirhosseinf79/renthub_service/internal/dto/jajiga"
	mihmansho_dto "github.com/amirhosseinf79/renthub_service/internal/dto/mihmansho"
	otaghak_dto "github.com/amirhosseinf79/renthub_service/internal/dto/otaghak"
	"github.com/amirhosseinf79/renthub_service/internal/services/requests"
	"github.com/google/uuid"
)

func (h *homsaService) generateEasyLoginBody(fields dto.ApiEasyLogin) any {
	switch h.service {
	case "homsa":
		return homsa_dto.HomsaLoginUserPass{
			Mobile:   fields.Username,
			Password: fields.Password,
			UseOTP:   false,
		}
	case "jajiga":
		return jajiga_dto.JajigaAuthRequestBody{
			Mobile:     fields.Username,
			Password:   &fields.Password,
			ISO2:       "IR",
			ClientID:   uuid.New().String(),
			ClientType: "browser",
		}
	case "otaghak":
		return otaghak_dto.OtaghakAuthRequestBody{
			UserName:     fields.Username,
			Password:     fields.Password,
			ClientId:     "Otaghak",
			ClientSecret: "secret",
			ArcValues:    map[string]string{},
		}
	case "mihmansho":
		return mihmansho_dto.AuthBody{
			Username: fields.Username,
			Password: fields.Password,
		}
	}
	return nil
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

func (h *homsaService) EasyLogin(fields dto.ApiEasyLogin) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID)
	endpoint := h.getEndpoints().LoginWithPass
	url, err := h.getFullURL(endpoint)
	if err != nil {
		log.FinalResult = err.Error()
		return log, err
	}
	header := h.getHeader()
	bodyRow := h.generateEasyLoginBody(fields)
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
	err = h.updateOrCreateAuthRecord(fields, response.GetToken())
	if err != nil {
		log.FinalResult = err.Error()
		return log, err
	}
	log.FinalResult = "success"
	log.IsSucceed = true
	return log, err
}
