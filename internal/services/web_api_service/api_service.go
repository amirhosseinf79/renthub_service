package cloner

import (
	"errors"
	"fmt"
	"sort"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/domain/repository"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	homsa_dto "github.com/amirhosseinf79/renthub_service/internal/dto/homsa"
	mihmansho_dto "github.com/amirhosseinf79/renthub_service/internal/dto/mihmansho"
)

type homsaService struct {
	logRepo     repository.LogRepository
	apiAuthRepo repository.ApiAuthRepository
	service     string
	apiSettings map[string]dto.ApiSettings
}

func NewHomsaService(apiAuthRepo repository.ApiAuthRepository, logRepo repository.LogRepository) interfaces.ApiService {
	return &homsaService{
		apiAuthRepo: apiAuthRepo,
		logRepo:     logRepo,
		apiSettings: map[string]dto.ApiSettings{
			"homsa": {
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
			"jabama": {
				ApiURL: "https://gw.jabama.com/mobile/api",
				Endpoints: dto.ApiEndpoints{
					LoginFirstStep:  dto.EndP{Address: "/v4/account/send-code", Method: "POST", ContentType: "body"},
					LoginSecondStep: dto.EndP{Address: "/v4/account/validate-code", Method: "POST", ContentType: "body"},
					GetProfile:      dto.EndP{Address: "/v1/profile?isHost=true", Method: "GET", ContentType: "body"},
					OpenCalendar:    dto.EndP{Address: "/v1/accommodations/host/Price/%v/price/calendar/enable", Method: "PUT", ContentType: "body"},
					CloseCalendar:   dto.EndP{Address: "/v1/accommodations/host/Price/%v/price/calendar/disable", Method: "PUT", ContentType: "body"},
					EditPricePerDay: dto.EndP{Address: "/taraaz/v1/pricing/management/accommodation/%v", Method: "PUT", ContentType: "body"},
				},
				Headers: map[string]string{
					"user-Agent":      "okhttp/4.12.0",
					"Accept":          "application/json",
					"accept-Encoding": "chunked",
					"Host":            "gw.jabama.com",
					"ab-channel":      "HostAndroid,3.6.9 - CafeBazaar,Android,%v",
					"Content-Type":    "application/json; charset=UTF-8",
					"Connection":      "Keep-Alive",
					"Authorization":   "Bearer %v",
				},
			},
			"jajiga": {
				ApiURL: "https://api.jajiga.com/api",
				Endpoints: dto.ApiEndpoints{
					LoginFirstStep:  dto.EndP{Address: "/auth/otp/send", Method: "POST", ContentType: "body"},
					LoginSecondStep: dto.EndP{Address: "/auth/login", Method: "POST", ContentType: "body"},
					LoginWithPass:   dto.EndP{Address: "/auth/logi", Method: "POST", ContentType: "body"},
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
					"content-type":    "application/json; charset=UTF-8",
					"accept-charset":  "UTF-8",
					"lang":            "fa",
					"Authorization":   "Bearer %v",
				},
			},
			"otaghak": {
				ApiURL: "https://core.otaghak.com",
				Endpoints: dto.ApiEndpoints{
					LoginFirstStep:  dto.EndP{Address: "/odata/Otaghak/Users/SendVerificationCode", Method: "POST", ContentType: "body"},
					LoginSecondStep: dto.EndP{Address: "/odata/Otaghak/Users/CheckVerificationCode", Method: "POST", ContentType: "body"},
					LoginWithPass:   dto.EndP{Address: "/api/v1/Identity/Login", Method: "POST", ContentType: "body"},
					GetProfile:      dto.EndP{Address: "/odata/Otaghak/Users/UserInfo", Method: "GET", ContentType: "body"},
					OpenCalendar:    dto.EndP{Address: "/odata/Otaghak/RoomBlockedUnblockedDays/ChangeBlockedDaysByHost", Method: "POST", ContentType: "body"},
					CloseCalendar:   dto.EndP{Address: "/odata/Otaghak/RoomBlockedUnblockedDays/ChangeBlockedDaysByHost", Method: "POST", ContentType: "body"},
					EditPricePerDay: dto.EndP{Address: "/odata/Otaghak/RoomPrices/GetHostShare", Method: "POST", ContentType: "body"},
					AddDiscount:     dto.EndP{Address: "/api/v1/HostDiscounts/CreateHostRoomDiscount", Method: "POST", ContentType: "body"},
					RemoveDiscount:  dto.EndP{Address: "/api/v1/HostDiscounts/ChangeHostRoomDiscountStatus", Method: "POST", ContentType: "body"},
					SetMinNight:     dto.EndP{Address: "/api/v1/HostDiscounts/ChangeHostRoomReserveRestriction", Method: "POST", ContentType: "body"},
					UnsetMinNight:   dto.EndP{Address: "/api/v1/HostDiscounts/ChangeHostRoomReserveRestriction", Method: "POST", ContentType: "body"},
				},
				Headers: map[string]string{
					"user-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:138.0) Gecko/20100101 Firefox/138.0",
					"Accept":          "application/json",
					"Accept-Encoding": "chunked",
					"Origin":          "https://www.otaghak.com",
					"host":            "core.otaghak.com",
					"Content-Type":    "application/json; charset=UTF-8",
					"accept-charset":  "UTF-8",
					"lang":            "fa",
					"Authorization":   "Bearer %v",
				},
			},
			"shab": {
				ApiURL: "https://api.shab.travel/api/fa/sandbox/v_1_4",
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
				},
				Headers: map[string]string{
					"user-Agent":      "Dart/2.19 (dart:io)",
					"accept":          "application/json",
					"accept-Encoding": "chunked",
					"Origin":          "https://www.shab.travel",
					"content-type":    "application/json; charset=UTF-8",
					"accept-charset":  "UTF-8",
					"lang":            "fa",
					"Authorization":   "Bearer %v",
				},
			},
		},
	}
}

func (h *homsaService) Set(service string) interfaces.ApiService {
	h.service = service
	return h
}

func (h *homsaService) RecordLog(log *models.Log) error {
	return h.logRepo.Create(log)
}

func (h *homsaService) initLog(userID uint, clientID string) *models.Log {
	return &models.Log{
		UserID:   userID,
		ClientID: clientID,
		Service:  h.service,
	}
}

func (h *homsaService) getFullURL(endpoint dto.EndP, vals ...any) (url string, err error) {
	errMsg := errors.New("service can not perform this action")
	if endpoint.Address == "" {
		err = errMsg
		return
	}
	realEndpoint := fmt.Sprintf(endpoint.Address, vals...)
	settings, ok := h.apiSettings[h.service]
	if !ok {
		err = errMsg
		return
	}
	url = settings.ApiURL + realEndpoint
	return
}

func (h *homsaService) getEndpoints() dto.ApiEndpoints {
	settings, ok := h.apiSettings[h.service]
	if !ok {
		return dto.ApiEndpoints{}
	}
	return settings.Endpoints
}

func (h *homsaService) getHeader() map[string]string {
	settings, ok := h.apiSettings[h.service]
	if !ok {
		return map[string]string{}
	}
	return settings.Headers
}

func (h *homsaService) getExtraHeader(token *models.ApiAuth) map[string]string {
	if h.service == "homsa" {
		return map[string]string{
			"authorization": token.AccessToken,
		}
	}
	return map[string]string{}
}

func (h *homsaService) generateEasyLoginBody(fields dto.ApiEasyLogin) any {
	if h.service == "homsa" {
		return homsa_dto.HomsaLoginUserPass{
			Mobile:   fields.Username,
			Password: fields.Password,
			UseOTP:   false,
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

func (h *homsaService) generateVerifyOTPBody(phoneNumber string, code string) any {
	if h.service == "homsa" {
		return homsa_dto.HomsaLoginUserPass{
			Mobile:   phoneNumber,
			Password: code,
			UseOTP:   true,
		}
	}
	return nil
}

func (h *homsaService) generateCalendarBody(roomID string, setOpen bool, dates []string) any {
	if h.service == "homsa" {
		if len(dates) > 1 {
			sort.Strings(dates)
		}
		return homsa_dto.HomsaCalendarBody{
			StartDate: dates[0],
			EndDate:   dates[len(dates)-1],
		}
	}
	return nil
}

func (h *homsaService) generatePriceBody(roomID string, amount int, dates []string) any {
	if h.service == "homsa" {
		if len(dates) > 1 {
			sort.Strings(dates)
		}
		return homsa_dto.HomsaPriceBody{
			StartDate:    dates[0],
			EndDate:      dates[len(dates)-1],
			Price:        amount,
			KeepDiscount: 0,
		}
	}
	return nil
}

func (h *homsaService) generateAddDiscountBody(roomID string, amount int, dates []string) any {
	if h.service == "homsa" {
		if len(dates) > 1 {
			sort.Strings(dates)
		}
		return homsa_dto.HomsaAddDiscountBody{
			StartDate:    dates[0],
			EndDate:      dates[len(dates)-1],
			Discount:     amount,
			KeepDiscount: 0,
		}
	}
	return nil
}

func (h *homsaService) generateRemoveDiscountBody(roomID string, dates []string) any {
	if h.service == "homsa" {
		if len(dates) > 1 {
			sort.Strings(dates)
		}
		return homsa_dto.HomsaRemoveDiscountBody{
			StartDate:    dates[0],
			EndDate:      dates[len(dates)-1],
			KeepDiscount: 0,
		}
	}
	return nil
}

func (h *homsaService) generateSetMinNightBody(roomID string, amount int, dates []string) any {
	if h.service == "homsa" {
		if len(dates) > 1 {
			sort.Strings(dates)
		}
		return homsa_dto.HomsaSetMinNightBody{
			StartDate: dates[0],
			EndDate:   dates[len(dates)-1],
			Min:       amount,
			Max:       nil,
		}
	}
	return nil
}

func (h *homsaService) generateUnsetMinNightBody(roomID string, dates []string) any {
	if h.service == "homsa" {
		if len(dates) > 1 {
			sort.Strings(dates)
		}
		return homsa_dto.HomsaUnsetMinNightBody{
			StartDate: dates[0],
			EndDate:   dates[len(dates)-1],
		}
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

func (h *homsaService) generateMihmanshoErrResponse() interfaces.ApiResponseManager {
	if h.service == "mihmansho" {
		return &mihmansho_dto.MihmanshoErrorResponse{}
	}
	return nil
}

func (h *homsaService) generateErrResponse() interfaces.ApiResponseManager {
	if h.service == "homsa" {
		return &homsa_dto.HomsaErrorResponse{}
	} else if h.service == "mihmansho" {
		return &mihmansho_dto.MihmanshoErrorResponse{}
	}
	return nil
}
