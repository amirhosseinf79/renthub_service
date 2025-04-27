package cloner

import (
	"errors"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/domain/repository"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	homsa_dto "github.com/amirhosseinf79/renthub_service/internal/dto/homsa"
	mihmansho_dto "github.com/amirhosseinf79/renthub_service/internal/dto/mihmansho"
)

type homsaService struct {
	apiAuthRepo repository.ApiAuthRepository
	service     string
	apiSettings map[string]dto.ApiSettings
}

func NewHomsaService(apiAuthRepo repository.ApiAuthRepository, service string) interfaces.ApiService {
	return &homsaService{
		apiAuthRepo: apiAuthRepo,
		service:     service,
		apiSettings: map[string]dto.ApiSettings{
			"homsa": {
				ApiURL: "https://www.homsa.net/api/v2",
				Endpoints: dto.ApiEndpoints{
					LoginFirstStep:  "/newAuth/otp/send",
					LoginSecondStep: "/newAuth/login",
					LoginWithPass:   "/newAuth/login",
					GetProfile:      "/host/profile",
					OpenCalendar:    "/host/room/%v/calendar/unblock",
					CloseCalendar:   "/host/room/%v/calendar/block",
					EditPricePerDay: "/host/room/%v/calendar/set_price",
					AddDiscount:     "/host/room/%v/calendar/set_discount",
					RemoveDiscount:  "/host/room/%v/calendar/remove_discount",
					SetMinNight:     "/host/room/%v/calendar/update_availability_rule",
					UnsetMinNight:   "/host/room/%v/calendar/delete_availability_rule",
				},
				Headers: map[string]string{
					"user-Agent":      "Dart/2.19 (dart:io)",
					"accept":          "application/json",
					"accept-Encoding": "chunked",
					"host":            "www.homsa.net",
					"content-type":    "application/json; charset=UTF-8",
					"accept-charset":  "UTF-8",
					"lang":            "fa",
					"authorization":   "bearer %v",
				},
			},
		},
	}
}

func (h *homsaService) getFullURL(endpoint string) (url string, err error) {
	if endpoint == "" {
		err = errors.New("service can not perform this action")
		return
	}
	url = h.apiSettings[h.service].ApiURL + endpoint
	return
}

func (h *homsaService) getEndpoints() dto.ApiEndpoints {
	return h.apiSettings[h.service].Endpoints
}

func (h *homsaService) getHeader() map[string]string {
	return h.apiSettings[h.service].Headers
}

func (h *homsaService) getExtraHeader(token *models.ApiAuth) map[string]string {
	if h.service == "homsa" {
		return map[string]string{
			"authorization": token.AccessToken,
		}
	}
	return map[string]string{}
}

func (h *homsaService) generateEasyLoginBody(fields dto.ApiEasyLogin, otp bool) any {
	if h.service == "homsa" {
		return homsa_dto.HomsaLoginUserPass{
			Mobile:   fields.Username,
			Password: fields.Password,
			UseOTP:   otp,
		}
	}
	return nil
}

func (h *homsaService) generateSendOTPBody(phoneNumber string) any {
	if h.service == "homsa" {
		return homsa_dto.HomsaOTPLogin{Mobile: phoneNumber}
	}
	return nil
}

func (h *homsaService) generateAuthResponse() interfaces.ApiResponseManager {
	if h.service == "homsa" {
		return &homsa_dto.HomsaAuthResponse{}
	}
	return nil
}

func (h *homsaService) generateOTPResponse() interfaces.ApiResponseManager {
	if h.service == "homsa" {
		return &homsa_dto.HomsaOTPResponse{}
	}
	return nil
}

func (h *homsaService) generateProfileResponse() interfaces.ApiResponseManager {
	if h.service == "mihmansho" {
		return &mihmansho_dto.MihmanshoProfileResponse{}
	}
	return nil
}

func (h *homsaService) generateErrResponse() interfaces.ApiResponseManager {
	if h.service == "homsa" {
		return &homsa_dto.HomsaErrorResponse{}
	}
	return nil
}
