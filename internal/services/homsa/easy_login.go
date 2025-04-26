package homsa

import (
	"errors"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/internal/services/requests"
)

func (h *homsaService) EasyLogin(fields dto.ApiEasyLogin) (err error) {
	if fields.Username == "" || fields.Password == "" {
		return
	}

	bodyRow := dto.HomsaLoginUserPass{
		Mobile:   fields.Username,
		Password: fields.Password,
		UseOTP:   false,
	}

	url := h.apiUrl + h.endpoints.LoginWithPass
	request := requests.New("POST", url, h.GetHeader(), map[string]string{})
	err = request.RequestBody(bodyRow)
	if err != nil {
		return err
	}
	request.PrintRequestDump()
	err = request.CommitRequest()
	if err != nil {
		return err
	}

	if !request.Ok() {
		var errResponse dto.HomsaErrorResponse
		err = request.Json(&errResponse)
		if err != nil {
			return err
		}
		err = errors.New(errResponse.Code)
		return err
	}

	var response dto.HomsaAuthResponse
	err = request.Json(&response)
	if err != nil {
		return err
	}

	exists := h.apiAuthRepo.CheckExists(fields.UserID, fields.ClientID, h.service)
	if exists {
		apiM, err := h.apiAuthRepo.GetByUnique(fields.UserID, fields.ClientID, h.service)
		if err != nil {
			return err
		}
		apiM.AccessToken = response.AccessToken
		apiM.RefreshToken = response.RefreshToken
		apiM.Username = fields.Username
		apiM.Password = fields.Password
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
			AccessToken:  response.AccessToken,
			RefreshToken: response.RefreshToken,
		}
		err = h.apiAuthRepo.Create(model)
		if err != nil {
			return err
		}
	}
	return nil
}
