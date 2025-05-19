package homsa

import (
	"bytes"
	"fmt"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	homsa_dto "github.com/amirhosseinf79/renthub_service/internal/dto/homsa"
)

type service struct {
	apiAuthService interfaces.ApiAuthInterface
	service        string
	apiSettings    dto.ApiSettings
	request        interfaces.FetchService
}

func New(apiAuthService interfaces.ApiAuthInterface, request interfaces.FetchService) interfaces.ApiService {
	return &service{
		service:        "homsa",
		apiAuthService: apiAuthService,
		request:        request,
		apiSettings: dto.ApiSettings{
			ApiURL: "https://www.homsa.net/api/v2",
			Endpoints: dto.ApiEndpoints{
				LoginFirstStep:  dto.EndP{Address: "/newAuth/otp/send", Method: "POST", ContentType: "body"},
				LoginSecondStep: dto.EndP{Address: "/newAuth/login", Method: "POST", ContentType: "body"},
				LoginWithPass:   dto.EndP{Address: "/newAuth/login", Method: "POST", ContentType: "body"},
				GetProfile:      dto.EndP{Address: "/host/profile", Method: "GET", ContentType: "body"},
				OpenCalendar:    dto.EndP{Address: "/host/room/%v/calendar/unblock", Method: "POST", ContentType: "body"},
				CloseCalendar:   dto.EndP{Address: "/host/room/%v/calendar/block", Method: "POST", ContentType: "body"},
				EditPricePerDay: dto.EndP{Address: "/host/room/%v/calendar/set_price", Method: "POST", ContentType: "body"},
				AddDiscount:     dto.EndP{Address: "/host/room/%v/calendar/set_discount", Method: "POST", ContentType: "body"},
				RemoveDiscount:  dto.EndP{Address: "/host/room/%v/calendar/remove_discount", Method: "POST", ContentType: "body"},
				SetMinNight:     dto.EndP{Address: "/host/room/%v/calendar/update_availability_rule", Method: "POST", ContentType: "body"},
				UnsetMinNight:   dto.EndP{Address: "/host/room/%v/calendar/delete_availability_rule", Method: "POST", ContentType: "body"},
			},
			Headers: map[string]string{
				"user-Agent":     "Dart/2.19 (dart:io)",
				"accept":         "application/json",
				"host":           "www.homsa.net",
				"accept-charset": "UTF-8",
				"lang":           "fa",
				"authorization":  "bearer %v",
			},
		},
	}
}

func (h *service) initLog(userID uint, clientID string) *models.Log {
	return &models.Log{
		UserID:   userID,
		ClientID: clientID,
		Service:  h.service,
	}
}

func (h *service) getFullURL(endpoint dto.EndP, vals ...any) (url string, err error) {
	errMsg := dto.ErrServiceUnavailable
	if endpoint.Address == "" {
		err = errMsg
		return
	}
	realEndpoint := endpoint.Address
	if bytes.Contains([]byte(endpoint.Address), []byte("%v")) {
		count := bytes.Count([]byte(endpoint.Address), []byte("%v"))
		if len(vals) >= count {
			realEndpoint = fmt.Sprintf(endpoint.Address, vals[:count]...)
		} else {
			return
		}
	}
	settings := h.apiSettings
	url = settings.ApiURL + realEndpoint
	return
}

func (h *service) getEndpoints() dto.ApiEndpoints {
	settings := h.apiSettings
	return settings.Endpoints
}

func (h *service) getHeader() map[string]string {
	settings := h.apiSettings
	return settings.Headers
}

func (h *service) getExtraHeader(token *models.ApiAuth) map[string]string {
	return map[string]string{
		"authorization": token.AccessToken,
	}
}

func (h *service) generateAuthResponse() interfaces.ApiResponseManager {
	return &homsa_dto.HomsaAuthResponse{}
}

func (h *service) generateOTPResponse() interfaces.ApiResponseManager {
	return &homsa_dto.HomsaOTPResponse{}
}

func (h *service) generateUpdateErrResponse() interfaces.ApiResponseManager {
	return &homsa_dto.HomsaErrorResponse{}
}

func (h *service) generateErrResponse() interfaces.ApiResponseManager {
	return &homsa_dto.HomsaErrorResponse{}
}

// Body
func (h *service) generateEasyLoginBody(fields dto.ApiEasyLogin) homsa_dto.HomsaLoginUserPass {
	return homsa_dto.HomsaLoginUserPass{
		Mobile:   fields.Username,
		Password: fields.Password,
		UseOTP:   false,
	}
}

func (h *service) generateSendOTPBody(phoneNumber string) homsa_dto.HomsaOTPLogin {
	return homsa_dto.HomsaOTPLogin{
		Mobile: phoneNumber,
	}
}

func (h *service) generateVerifyOTPBody(phoneNumber string, code string) homsa_dto.HomsaLoginUserPass {
	return homsa_dto.HomsaLoginUserPass{
		Mobile:   phoneNumber,
		Password: code,
		UseOTP:   true,
	}
}

func (h *service) generateCalendarBody(dates []string) any {
	return homsa_dto.HomsaCalendarBody{
		StartDate: dates[0],
		EndDate:   dates[len(dates)-1],
	}
}

func (h *service) generateSetMinNightBody(amount int, dates []string) homsa_dto.HomsaSetMinNightBody {
	return homsa_dto.HomsaSetMinNightBody{
		StartDate: dates[0],
		EndDate:   dates[len(dates)-1],
		Min:       amount,
		Max:       nil,
	}
}

func (h *service) generateUnsetMinNightBody(dates []string) homsa_dto.HomsaUnsetMinNightBody {
	return homsa_dto.HomsaUnsetMinNightBody{
		StartDate: dates[0],
		EndDate:   dates[len(dates)-1],
	}
}

func (h *service) generatePriceBody(amount int, dates []string) homsa_dto.HomsaPriceBody {
	return homsa_dto.HomsaPriceBody{
		StartDate:    dates[0],
		EndDate:      dates[len(dates)-1],
		Price:        amount,
		KeepDiscount: 0,
	}
}

func (h *service) generateAddDiscountBody(amount int, dates []string) homsa_dto.HomsaAddDiscountBody {
	return homsa_dto.HomsaAddDiscountBody{
		StartDate:    dates[0],
		EndDate:      dates[len(dates)-1],
		Discount:     amount,
		KeepDiscount: 0,
	}
}

func (h *service) generateRemoveDiscountBody(dates []string) homsa_dto.HomsaRemoveDiscountBody {
	return homsa_dto.HomsaRemoveDiscountBody{
		StartDate:    dates[0],
		EndDate:      dates[len(dates)-1],
		KeepDiscount: 0,
	}
}
