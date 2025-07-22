package jabama

import (
	"bytes"
	"fmt"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	jabama_dto "github.com/amirhosseinf79/renthub_service/internal/dto/jabama"
	"github.com/google/uuid"
)

type service struct {
	apiAuthService interfaces.ApiAuthInterface
	service        string
	apiSettings    dto.ApiSettings
	request        interfaces.FetchService
}

func New(apiAuthService interfaces.ApiAuthInterface, request interfaces.FetchService) interfaces.ApiService {
	return &service{
		service:        "jabama",
		apiAuthService: apiAuthService,
		request:        request,
		apiSettings: dto.ApiSettings{
			ApiURL: "https://gw.jabama.com/mobile/api",
			Endpoints: dto.ApiEndpoints{
				LoginFirstStep:     dto.EndP{Address: "/v4/account/send-code", Method: "POST", ContentType: "body"},
				LoginSecondStep:    dto.EndP{Address: "/v4/account/validate-code", Method: "POST", ContentType: "body"},
				LoginWithPass:      dto.EndP{Address: "/v4/account/login/refresh/%v", Method: "POST", ContentType: "body"},
				GetProfile:         dto.EndP{Address: "/v1/profile?isHost=true", Method: "GET", ContentType: "body"},
				OpenCalendar:       dto.EndP{Address: "/v1/accommodations/host/Price/%v/price/calendar/enable", Method: "PUT", ContentType: "body"},
				CloseCalendar:      dto.EndP{Address: "/v1/accommodations/host/Price/%v/price/calendar/disable", Method: "PUT", ContentType: "body"},
				EditPricePerDay:    dto.EndP{Address: "/taraaz/v1/pricing/management/accommodation/%v", Method: "PUT", ContentType: "body"},
				AddDiscount:        dto.EndP{Address: "/taraaz/v1/accommodation/host/base-price/update/calendar/%v", Method: "PUT", ContentType: "body"},
				RemoveDiscount:     dto.EndP{Address: "/taraaz/v1/accommodation/host/base-price/update/calendar/%v", Method: "PUT", ContentType: "body"},
				GetCalendarDetails: dto.EndP{Address: "/taraaz/v1/pricing/management/accommodation/%v", Method: "POST", ContentType: "body"},
				GETRooms:           dto.EndP{Address: "/v2/accommodation/host/accommodation", Method: "GET", ContentType: "query"},
				// GETRoomDetails:  dto.EndP{Address: "/v2/accommodation/host/accommodation", Method: "GET", ContentType: "query"},
			},
			Headers: map[string]string{
				"Authorization":   "Bearer %v",
				"Accept":          "application/json",
				"Accept-Charset":  "UTF-8",
				"User-Agent":      "okhttp/4.12.0",
				"Host":            "gw.jabama.com",
				"Connection":      "Keep-Alive",
				"Accept-Encoding": "gzip",
				"ab-channel":      "HostAndroid,3.6.9 - CafeBazaar,Android,%v",
			},
		},
	}
}

func (h *service) initLog(userID uint, clientID string, action string) *models.Log {
	return &models.Log{
		UserID:   userID,
		ClientID: clientID,
		Service:  h.service,
		Action:   action,
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
		"ab-channel":    uuid.New().String(),
	}

}

func (h *service) generateAuthResponse() interfaces.ApiResponseManager {
	return &jabama_dto.Response{}
}

func (h *service) generateOTPResponse() interfaces.ApiResponseManager {
	return &jabama_dto.Response{}
}

func (h *service) generateUpdateErrResponse() interfaces.ApiResponseManager {
	return &jabama_dto.UpdateErrorResponse{}
}

func (h *service) generateErrResponse() interfaces.ApiResponseManager {
	return &jabama_dto.Response{}
}

// Body
func (h *service) generateCalendarBody(dates []string) jabama_dto.OpenClosCalendar {
	return jabama_dto.OpenClosCalendar{
		Dates: dates,
	}
}

func (h *service) generateSendOTPBody(phoneNumber string) jabama_dto.OTPLogin {
	return jabama_dto.OTPLogin{
		Mobile: phoneNumber,
	}
}

func (h *service) generateVerifyOTPBody(phoneNumber string, code string) jabama_dto.OTPLogin {
	return jabama_dto.OTPLogin{
		Mobile: phoneNumber,
		Code:   code,
	}
}

func (h *service) generatePriceBody(fields *dto.UpdateFields) jabama_dto.EditPricePerDay {
	return jabama_dto.EditPricePerDay{
		Type:  nil,
		Days:  fields.Dates,
		Value: fields.Amount * 10,
	}
}

func (h *service) generateDiscountBody(fields *dto.UpdateFields, currentPrice int) jabama_dto.DiscountDTO {
	return jabama_dto.DiscountDTO{
		Days:     fields.Dates,
		Discount: fields.Amount,
		Price:    currentPrice,
	}
}

func (h *service) generateCalendarDetailsBody(days []string) jabama_dto.CalendarBody {
	return jabama_dto.CalendarBody{Days: days}
}
