package cloner

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/domain/repository"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	homsa_dto "github.com/amirhosseinf79/renthub_service/internal/dto/homsa"
	jabama_dto "github.com/amirhosseinf79/renthub_service/internal/dto/jabama"
	jajiga_dto "github.com/amirhosseinf79/renthub_service/internal/dto/jajiga"
	mihmansho_dto "github.com/amirhosseinf79/renthub_service/internal/dto/mihmansho"
	otaghak_dto "github.com/amirhosseinf79/renthub_service/internal/dto/otaghak"
	shab_dto "github.com/amirhosseinf79/renthub_service/internal/dto/shab"
	"github.com/google/uuid"
	ptime "github.com/yaa110/go-persian-calendar"
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
					"user-Agent":     "Dart/2.19 (dart:io)",
					"accept":         "application/json",
					"host":           "www.homsa.net",
					"content-type":   "application/json; charset=UTF-8",
					"accept-charset": "UTF-8",
					"lang":           "fa",
					"authorization":  "bearer %v",
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
					"Authorization":   "Bearer %v",
					"Accept":          "application/json",
					"Accept-Charset":  "UTF-8",
					"User-Agent":      "okhttp/4.12.0",
					"Content-Type":    "application/json; charset=UTF-8",
					"Host":            "gw.jabama.com",
					"Connection":      "Keep-Alive",
					"Accept-Encoding": "gzip",
					"ab-channel":      "HostAndroid,3.6.9 - CafeBazaar,Android,%v",
				},
			},
			"jajiga": {
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
					"content-type":    "application/json; charset=UTF-8",
					"accept-charset":  "UTF-8",
					"lang":            "fa",
					"Authorization":   "Bearer %v",
				},
			},
			"mihmansho": {
				ApiURL: "https://www.mihmansho.com/myapi/v1",
				Endpoints: dto.ApiEndpoints{
					LoginFirstStep:  dto.EndP{Address: "/loginfirststep", Method: "POST", ContentType: "query"},
					LoginSecondStep: dto.EndP{Address: "/loginwithcode", Method: "POST", ContentType: "query"},
					LoginWithPass:   dto.EndP{Address: "/login", Method: "POST", ContentType: "query"},
					GetProfile:      dto.EndP{Address: "/getprofile", Method: "GET", ContentType: "body"},
					OpenCalendar:    dto.EndP{Address: "/ReserveDates", Method: "PUT", ContentType: "formData"},
					CloseCalendar:   dto.EndP{Address: "/ReserveDates", Method: "PUT", ContentType: "formData"},
					EditPricePerDay: dto.EndP{Address: "/EditPricesCalendar?ProductId=%v&Price=%v&AddedGuestPrice=0", Method: "PUT", ContentType: "unlencoded"},
				},
				Headers: map[string]string{
					"user-agent": "okhttp/3.12.1",
					"version":    "1.3.6",
					"city":       "0",
					"ucode":      "%v",
					"token":      "%v",
					"imei":       "%v",
				},
			},
			"otaghak": {
				ApiURL: "https://core.otaghak.com",
				Endpoints: dto.ApiEndpoints{
					LoginFirstStep:  dto.EndP{Address: "/odata/Otaghak/Users/SendVerificationCode", Method: "POST", ContentType: "body"},
					LoginSecondStep: dto.EndP{Address: "/api/v1/Identity/Login", Method: "POST", ContentType: "body"},
					LoginWithPass:   dto.EndP{Address: "/api/v1/Identity/Login", Method: "POST", ContentType: "body"},
					GetProfile:      dto.EndP{Address: "/odata/Otaghak/Users/UserInfo", Method: "GET", ContentType: "body"},
					OpenCalendar:    dto.EndP{Address: "/odata/Otaghak/RoomBlockedUnblockedDays/ChangeBlockedDaysByHost", Method: "POST", ContentType: "body"},
					CloseCalendar:   dto.EndP{Address: "/odata/Otaghak/RoomBlockedUnblockedDays/ChangeBlockedDaysByHost", Method: "POST", ContentType: "body"},
					EditPricePerDay: dto.EndP{Address: "/odata/Otaghak/RoomPerDayPrice/UpdateRoomPerDayPriceByHost", Method: "POST", ContentType: "body"},
					AddDiscount:     dto.EndP{Address: "/api/v1/HostDiscounts/CreateHostRoomDiscount", Method: "POST", ContentType: "body"},
					RemoveDiscount:  dto.EndP{Address: "/api/v1/HostDiscounts/ChangeHostRoomDiscountStatus", Method: "POST", ContentType: "body"},
					SetMinNight:     dto.EndP{Address: "/api/v1/HostDiscounts/ChangeHostRoomReserveRestriction", Method: "POST", ContentType: "body"},
					UnsetMinNight:   dto.EndP{Address: "/api/v1/HostDiscounts/ChangeHostRoomReserveRestriction", Method: "POST", ContentType: "body"},
				},
				Headers: map[string]string{
					"user-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:138.0) Gecko/20100101 Firefox/138.0",
					"Accept":          "application/json, text/plain, */*",
					"Accept-Encoding": "chunked",
					"Origin":          "https://www.otaghak.com",
					"Host":            "core.otaghak.com",
					"Content-Type":    "application/json;charset=UTF-8",
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
					"User-Agent":    "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:138.0) Gecko/20100101 Firefox/138.0",
					"Accept":        "application/json, text/plain, */*",
					"Content-Type":  "application/json;charset=UTF-8",
					"Connection":    "keep-alive",
					"Authorization": "Bearer %v",
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

func (h *homsaService) datesToIso(dates []string) []string {
	var formattedDates []string
	for _, item := range dates {
		formattedDates = append(formattedDates, fmt.Sprintf("%vT00:00:00.000Z", item))
	}
	return formattedDates
}

func (h *homsaService) datesToJalali(dates []string) []string {
	var jdates []string
	for _, date := range dates {
		parsedTime, err := time.Parse("2006-01-02", date)

		if err != nil {
			return nil
		}

		// Now convert to Jalali
		ptobj := ptime.New(parsedTime)

		// Format it as jYY-jMM-jDD
		jalaliDate := ptobj.Format("yyyy-MM-dd")
		jdates = append(jdates, jalaliDate)
		fmt.Println(jdates)
	}
	return jdates
}

func (h *homsaService) getFullURL(endpoint dto.EndP, vals ...any) (url string, err error) {
	errMsg := errors.New("service can not perform this action")
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

	switch h.service {
	case "homsa":
		return map[string]string{
			"authorization": token.AccessToken,
		}
	case "jabama":
		return map[string]string{
			"Authorization": token.AccessToken,
			"ab-channel":    uuid.New().String(),
		}
	case "mihmansho":
		return map[string]string{
			"ucode": token.Ucode,
			"token": token.AccessToken,
			"imei":  uuid.New().String(),
		}
	default:
		return map[string]string{
			"Authorization": token.AccessToken,
		}
	}
}

func (h *homsaService) generateAuthResponse() interfaces.ApiResponseManager {
	switch h.service {
	case "homsa":
		return &homsa_dto.HomsaAuthResponse{}
	case "jabama":
		return &jabama_dto.Response{}
	case "jajiga":
		return &jajiga_dto.AuthOkResponse{}
	case "otaghak":
		return &otaghak_dto.AuthOkResponse{}
	case "shab":
		return &shab_dto.AuthResponse{}
	default:
		return nil
	}
}

func (h *homsaService) generateOTPResponse() interfaces.ApiResponseManager {
	switch h.service {
	case "homsa":
		return &homsa_dto.HomsaOTPResponse{}
	case "jabama":
		return &jabama_dto.Response{}
	case "jajiga":
		return &jajiga_dto.OTPResponse{}
	case "otaghak":
		return &otaghak_dto.OTPResponse{}
	case "shab":
		return &shab_dto.AuthOTPResponse{}
	default:
		return nil
	}
}

func (h *homsaService) generateProfileResponse() interfaces.ApiResponseManager {
	switch h.service {
	case "mihmansho":
		return &mihmansho_dto.MihmanshoProfileResponse{}
	default:
		return nil
	}
}

func (h *homsaService) generateUpdateErrResponse() interfaces.ApiResponseManager {
	switch h.service {
	case "mihmansho":
		return &mihmansho_dto.MihmanshoErrorResponse{}
	case "homsa":
		return &homsa_dto.HomsaErrorResponse{}
	case "jabama":
		return &jabama_dto.UpdateErrorResponse{}
	case "jajiga":
		return &jajiga_dto.ErrorResponse{}
	case "otaghak":
		return &otaghak_dto.ErrorResponse{}
	case "shab":
		return &shab_dto.ErrResponse{}
	default:
		return nil
	}
}

func (h *homsaService) generateErrResponse() interfaces.ApiResponseManager {
	switch h.service {
	case "homsa":
		return &homsa_dto.HomsaErrorResponse{}
	case "mihmansho":
		return &mihmansho_dto.MihmanshoErrorResponse{}
	case "jabama":
		return &jabama_dto.Response{}
	case "jajiga":
		return &jajiga_dto.ErrorResponse{}
	case "otaghak":
		return &otaghak_dto.ErrorResponse{}
	case "shab":
		return &shab_dto.ErrResponse{}
	default:
		return nil
	}
}
