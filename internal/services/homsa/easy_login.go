package homsa

import (
	"encoding/json"
	"errors"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/internal/services/requests"
)

func (h *homsaService) EasyLogin(fields dto.ApiEasyLogin) (final *dto.TokenResponse, err error) {
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
		return
	}
	request.PrintRequestDump()
	resp, err := request.CommitRequest()
	if err != nil {
		return
	}

	if request.Ok() {
		var errResponse dto.HomsaErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&errResponse)
		if err != nil {
			return
		}
		err = errors.New(errResponse.Code)
		return
	}

	var response dto.HomsaAuthResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return
	}

	exists := h.apiAuthRepo.CheckExists(fields.UserID, fields.ClientID, h.service)
	if exists {
		apiM, err := h.apiAuthRepo.GetByUnique(fields.UserID, fields.ClientID, h.service)
		if err != nil {
			return nil, err
		}
		apiM.AccessToken = response.AccessToken
		apiM.RefreshToken = response.RefreshToken
		apiM.Username = fields.Username
		apiM.Password = fields.Password
		err = h.apiAuthRepo.Update(apiM)
		if err != nil {
			return nil, err
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
			return nil, err
		}
	}

	final = &dto.TokenResponse{
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
	}
	return final, nil
}
