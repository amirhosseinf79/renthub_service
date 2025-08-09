package shab

import (
	"bytes"
	"fmt"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	shab_dto "github.com/amirhosseinf79/renthub_service/internal/dto/shab"
	"github.com/amirhosseinf79/renthub_service/pkg"
)

type service struct {
	apiAuthService interfaces.ApiAuthInterface
	service        string
	apiSettings    dto.ApiSettings
	request        interfaces.FetchService
}

func New(apiAuthService interfaces.ApiAuthInterface, request interfaces.FetchService) interfaces.ApiService {
	return &service{
		service:        "shab",
		apiAuthService: apiAuthService,
		request:        request,
		apiSettings: dto.ApiSettings{
			ApiURL: "https://api.shab.ir/api/fa/sandbox/v_1_4",
			Endpoints: dto.ApiEndpoints{
				LoginFirstStep:  dto.EndP{Address: "/auth/login-otp", Method: "POST", ContentType: "body"},
				LoginSecondStep: dto.EndP{Address: "/auth/verify-otp", Method: "POST", ContentType: "body"},
				GetProfile:      dto.EndP{Address: "/bootstrap", Method: "GET", ContentType: "body"},
				OpenCalendar:    dto.EndP{Address: "/house/%v/calendar", Method: "POST", ContentType: "body"},
				CloseCalendar:   dto.EndP{Address: "/house/%v/calendar", Method: "POST", ContentType: "body"},
				EditPricePerDay: dto.EndP{Address: "/house/%v/calendar", Method: "POST", ContentType: "body"},
				AddDiscount:     dto.EndP{Address: "/house/%v/calendar", Method: "POST", ContentType: "body"},
				RemoveDiscount:  dto.EndP{Address: "/house/%v/calendar", Method: "POST", ContentType: "body"},
				SetMinNight:     dto.EndP{Address: "/house/%v/calendar", Method: "POST", ContentType: "body"},
				UnsetMinNight:   dto.EndP{Address: "/house/%v/calendar", Method: "POST", ContentType: "body"},
				GetReservations: dto.EndP{Address: "/reserve", Method: "GET", ContentType: "query"},
			},
			Headers: map[string]string{
				"User-Agent":    "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:138.0) Gecko/20100101 Firefox/138.0",
				"Accept":        "application/json, text/plain, */*",
				"Connection":    "keep-alive",
				"Authorization": "Bearer %v",
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
	}
}

func (h *service) generateAuthResponse() interfaces.ApiResponseManager {
	return &shab_dto.AuthResponse{}
}

func (h *service) generateOTPResponse() interfaces.ApiResponseManager {
	return &shab_dto.AuthOTPResponse{}
}

func (h *service) generateUpdateErrResponse() interfaces.ApiResponseManager {
	return &shab_dto.ErrResponse{}
}

func (h *service) generateErrResponse() interfaces.ApiResponseManager {
	return &shab_dto.ErrResponse{}
}

// Body
func (h *service) generateCalendarBody(setOpen bool, dates []string) shab_dto.CalendarBody {
	status := "set_disabled"
	if setOpen {
		status = "unset_disabled"
	}
	var jdates []string
	fmt.Println("shab", dates)
	tmpDates := pkg.DatesToJalali(dates, true)
	jdates = append(jdates, tmpDates...)
	return shab_dto.CalendarBody{
		Action: status,
		Dates:  jdates,
	}
}

func (h *service) generateAddDiscountBody(amount int, dates []string) shab_dto.EditDiscountBody {
	jdates := pkg.DatesToJalali(dates, true)
	return shab_dto.EditDiscountBody{
		Action:        "set_daily_discount",
		Dates:         jdates,
		DailyDiscount: amount,
	}
}

func (h *service) generateRemoveDiscountBody(dates []string) shab_dto.UnsetDiscountBody {
	jdates := pkg.DatesToJalali(dates, true)
	return shab_dto.UnsetDiscountBody{
		Action: "unset_daily_discount",
		Dates:  jdates,
	}
}

func (h *service) generateSendOTPBody(phoneNumber string) shab_dto.OTPBody {
	return shab_dto.OTPBody{
		Mobile:      phoneNumber,
		CountryCode: "+98",
	}
}

func (h *service) generateVerifyOTPBody(phoneNumber string, code string) shab_dto.VerifyOTOBody {
	return shab_dto.VerifyOTOBody{
		Mobile:      phoneNumber,
		CountryCode: "+98",
		Code:        code,
	}
}

func (h *service) generateSetMinNightBody(amount int, dates []string) shab_dto.EditMinNightBody {
	jdates := pkg.DatesToJalali(dates, true)
	return shab_dto.EditMinNightBody{
		Action:  "set_min_days",
		Dates:   jdates,
		MinDays: amount,
	}
}

func (h *service) generateUnsetMinNightBody(dates []string) shab_dto.EditMinNightBody {
	jdates := pkg.DatesToJalali(dates, true)
	return shab_dto.EditMinNightBody{
		Action:  "set_min_days",
		Dates:   jdates,
		MinDays: 1,
	}
}

func (h *service) generatePriceBody(amount int, dates []string) shab_dto.EditPriceBody {
	jdates := pkg.DatesToJalali(dates, true)
	return shab_dto.EditPriceBody{
		KeepDiscount: false,
		Price:        amount / 1000,
		Dates:        jdates,
	}
}
