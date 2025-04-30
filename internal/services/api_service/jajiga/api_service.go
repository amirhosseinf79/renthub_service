package jajiga

import (
	"bytes"
	"fmt"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	jajiga_dto "github.com/amirhosseinf79/renthub_service/internal/dto/jajiga"
	"github.com/google/uuid"
)

type service struct {
	apiAuthService interfaces.ApiAuthInterface
	service        string
	apiSettings    dto.ApiSettings
}

func New(apiAuthService interfaces.ApiAuthInterface) interfaces.ApiService {
	return &service{
		service:        "jajiga",
		apiAuthService: apiAuthService,
		apiSettings: dto.ApiSettings{
			ApiURL: "https://api.jajiga.com/api",
			Endpoints: dto.ApiEndpoints{
				LoginFirstStep:  dto.EndP{Address: "/auth/otp/send", Method: "POST", ContentType: "body"},
				LoginSecondStep: dto.EndP{Address: "/auth/login", Method: "POST", ContentType: "body"},
				LoginWithPass:   dto.EndP{Address: "/auth/login", Method: "POST", ContentType: "body"},
				GetProfile:      dto.EndP{Address: "/userinfo", Method: "GET", ContentType: "body"},
				OpenCalendar:    dto.EndP{Address: "/nights", Method: "PUT", ContentType: "body"},
				CloseCalendar:   dto.EndP{Address: "/nights", Method: "PUT", ContentType: "body"},
				EditPricePerDay: dto.EndP{Address: "/nights", Method: "PUT", ContentType: "body"},
				AddDiscount:     dto.EndP{Address: "/nights", Method: "PUT", ContentType: "body"},
				RemoveDiscount:  dto.EndP{Address: "/nights", Method: "PUT", ContentType: "body"},
				SetMinNight:     dto.EndP{Address: "/nights", Method: "PUT", ContentType: "body"},
				UnsetMinNight:   dto.EndP{Address: "/nights", Method: "PUT", ContentType: "body"},
			},
			Headers: map[string]string{
				"user-Agent":      "Dart/2.19 (dart:io)",
				"accept":          "application/json",
				"accept-Encoding": "chunked",
				"host":            "api.jajiga.com",
				"Host":            "api.jajiga.com",
				"accept-charset":  "UTF-8",
				"lang":            "fa",
				"Authorization":   "Bearer %v",
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
		"Authorization": token.AccessToken,
	}
}

func (h *service) generateAuthResponse() interfaces.ApiResponseManager {
	return &jajiga_dto.AuthOkResponse{}
}

func (h *service) generateOTPResponse() interfaces.ApiResponseManager {
	return &jajiga_dto.OTPResponse{}
}

func (h *service) generateProfileResponse() interfaces.ApiResponseManager {
	return nil
}

func (h *service) generateUpdateErrResponse() interfaces.ApiResponseManager {
	return &jajiga_dto.ErrorResponse{}
}

func (h *service) generateErrResponse() interfaces.ApiResponseManager {
	return &jajiga_dto.ErrorResponse{}
}

// Body
func (h *service) generateCalendarBody(roomID string, setOpen bool, dates []string) any {
	var num int
	if !setOpen {
		num = 1
	}
	return jajiga_dto.CalendarBody{
		RoomID:       roomID,
		Dates:        dates,
		DisableCount: num,
	}
}

func (h *service) generateAddDiscountBody(roomID string, amount int, dates []string) any {
	return jajiga_dto.DiscountBody{
		RoomID:  roomID,
		Dates:   dates,
		Percent: amount,
	}
}

func (h *service) generateRemoveDiscountBody(roomID string, dates []string) any {
	return jajiga_dto.DiscountBody{
		RoomID:  roomID,
		Dates:   dates,
		Percent: 0,
	}
}

func (h *service) generateEasyLoginBody(fields dto.ApiEasyLogin) any {
	return jajiga_dto.JajigaAuthRequestBody{
		Mobile:     fields.Username,
		Password:   &fields.Password,
		ISO2:       "IR",
		ClientID:   uuid.New().String(),
		ClientType: "browser",
	}
}

func (h *service) generateSendOTPBody(phoneNumber string) any {
	return jajiga_dto.OTPLogin{
		Mobile: phoneNumber,
		ISO2:   "IR",
	}
}

func (h *service) generateVerifyOTPBody(phoneNumber string, code string) any {
	return jajiga_dto.JajigaTokenAuthRequestBody{
		Mobile:   phoneNumber,
		Token:    &code,
		ClientID: uuid.New().String(),
		ISO2:     "IR",
	}
}

func (h *service) generateSetMinNightBody(roomID string, amount int, dates []string) any {
	return jajiga_dto.MinNightBody{
		RoomID:    roomID,
		Dates:     dates,
		MinNights: amount,
	}
}

func (h *service) generateUnsetMinNightBody(roomID string, dates []string) any {
	return jajiga_dto.MinNightBody{
		RoomID:    roomID,
		Dates:     dates,
		MinNights: 1,
	}
}

func (h *service) generatePriceBody(roomID string, amount int, dates []string) any {
	return jajiga_dto.PriceBody{
		RoomID: roomID,
		Dates:  dates,
		Price:  amount,
	}
}
