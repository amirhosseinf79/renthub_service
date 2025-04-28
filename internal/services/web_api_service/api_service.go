package cloner

import (
	"bytes"
	"errors"
	"fmt"
	"sort"
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
			"otaghak": {
				ApiURL: "https://core.otaghak.com",
				Endpoints: dto.ApiEndpoints{
					LoginFirstStep:  dto.EndP{Address: "/odata/Otaghak/Users/SendVerificationCode", Method: "POST", ContentType: "body"},
					LoginSecondStep: dto.EndP{Address: "/api/v1/Identity/Login", Method: "POST", ContentType: "body"},
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
					"content-type":    "application/json; charset=UTF-8",
					"accept-charset":  "UTF-8",
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
		realEndpoint = fmt.Sprintf(endpoint.Address, vals...)
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
	if h.service == "homsa" {
		return map[string]string{
			"authorization": token.AccessToken,
		}
	} else if h.service == "jabama" {
		return map[string]string{
			"Authorization": token.AccessToken,
			"ab-channel":    uuid.New().String(),
		}
	} else {
		return map[string]string{
			"Authorization": token.AccessToken,
		}
	}
}

func (h *homsaService) generateEasyLoginBody(fields dto.ApiEasyLogin) any {
	if h.service == "homsa" {
		return homsa_dto.HomsaLoginUserPass{
			Mobile:   fields.Username,
			Password: fields.Password,
			UseOTP:   false,
		}
	} else if h.service == "jajiga" {
		return jajiga_dto.JajigaAuthRequestBody{
			Mobile:     fields.Username,
			Password:   &fields.Password,
			ISO2:       "IR",
			ClientID:   uuid.New().String(),
			ClientType: "browser",
		}
	} else if h.service == "otaghak" {
		return otaghak_dto.OtaghakAuthRequestBody{
			UserName:     fields.Username,
			Password:     fields.Password,
			ClientId:     "Otaghak",
			ClientSecret: "secret",
			ArcValues:    map[string]string{},
		}
	}
	return nil
}

func (h *homsaService) generateSendOTPBody(phoneNumber string) any {
	if h.service == "homsa" {
		return homsa_dto.HomsaOTPLogin{Mobile: phoneNumber}
	} else if h.service == "jabama" {
		return jabama_dto.OTPLogin{Mobile: phoneNumber}
	} else if h.service == "jajiga" {
		return jajiga_dto.OTPLogin{
			Mobile: phoneNumber,
			ISO2:   "IR",
		}
	} else if h.service == "otaghak" {
		return otaghak_dto.OTPBody{
			UserName:   phoneNumber,
			IsShortOtp: true,
		}
	} else if h.service == "shab" {
		return shab_dto.OTPBody{
			Mobile:      phoneNumber,
			CountryCode: "+98",
		}
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
	} else if h.service == "jabama" {
		return jabama_dto.OTPLogin{
			Mobile: phoneNumber,
			Code:   code,
		}
	} else if h.service == "jajiga" {
		return jajiga_dto.JajigaTokenAuthRequestBody{
			Mobile:   phoneNumber,
			Token:    &code,
			ClientID: uuid.New().String(),
			ISO2:     "IR",
		}
	} else if h.service == "otaghak" {
		return otaghak_dto.OtaghakAuthRequestBody{
			UserName:     phoneNumber,
			ClientId:     "Otaghak",
			ClientSecret: "secret",
			ArcValues:    map[string]string{"OtpCode": code},
		}
	} else if h.service == "shab" {
		return shab_dto.VerifyOTOBody{
			Mobile:      phoneNumber,
			CountryCode: "+98",
			Code:        code,
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
	} else if h.service == "jabama" {
		return jabama_dto.OpenClosCalendar{
			Dates: dates,
		}
	} else if h.service == "jajiga" {
		var num int
		if !setOpen {
			num = 1
		}
		return jajiga_dto.CalendarBody{
			RoomID:       roomID,
			Dates:        dates,
			DisableCount: num,
		}
	} else if h.service == "otaghak" {
		if setOpen {
			return otaghak_dto.CalendarBody{
				RoomID:        roomID,
				UnblockedDays: h.datesToIso(dates),
			}
		}
		return otaghak_dto.CalendarBody{
			RoomID:      roomID,
			BlockedDays: h.datesToIso(dates),
		}
	} else if h.service == "shab" {
		status := "set_disabled"
		if setOpen {
			status = "unset_disabled"
		}
		return shab_dto.CalendarBody{
			Action: status,
			Dates:  h.datesToJalali(dates),
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
	} else if h.service == "jabama" {
		return jabama_dto.EditPricePerDay{
			Type:  nil,
			Days:  dates,
			Value: amount * 10,
		}
	} else if h.service == "jajiga" {
		return jajiga_dto.PriceBody{
			RoomID: roomID,
			Dates:  dates,
			Price:  amount,
		}
	} else if h.service == "otaghak" {
		formattedDates := h.datesToIso(dates)
		var formattedDays []otaghak_dto.DayPricePair
		for _, item := range formattedDates {
			formattedDays = append(formattedDays, otaghak_dto.DayPricePair{Day: item, Price: amount})
		}
		return otaghak_dto.EditPriceBody{
			RoomID:       roomID,
			PerDayPrices: formattedDays,
		}
	} else if h.service == "shab" {
		return shab_dto.EditPriceBody{
			KeepDiscount: false,
			Price:        amount,
			Dates:        h.datesToJalali(dates),
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
	} else if h.service == "jajiga" {
		return jajiga_dto.DiscountBody{
			RoomID:  roomID,
			Dates:   dates,
			Percent: amount,
		}
	} else if h.service == "otaghak" {
		return otaghak_dto.EditDiscountBody{
			DiscountPercent: amount,
			EffectiveDays:   h.datesToIso(dates),
			RoomID:          roomID,
		}
	} else if h.service == "shab" {
		return shab_dto.EditDiscountBody{
			Action:        "set_daily_discount",
			Dates:         h.datesToJalali(dates),
			DailyDiscount: amount,
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
	} else if h.service == "jajiga" {
		return jajiga_dto.DiscountBody{
			RoomID:  roomID,
			Dates:   dates,
			Percent: 0,
		}
	} else if h.service == "otaghak" {
		return otaghak_dto.EditDiscountBody{
			DiscountPercent: 0,
			EffectiveDays:   h.datesToIso(dates),
			RoomID:          roomID,
		}
	} else if h.service == "shab" {
		return shab_dto.EditDiscountBody{
			Action: "unset_daily_discount",
			Dates:  h.datesToJalali(dates),
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
	} else if h.service == "jajiga" {
		return jajiga_dto.MinNightBody{
			RoomID:    roomID,
			Dates:     dates,
			MinNights: amount,
		}
	} else if h.service == "otaghak" {
		return otaghak_dto.EditMinNightBody{
			MinNights:     amount,
			EffectiveDays: h.datesToIso(dates),
			RoomID:        roomID,
		}
	} else if h.service == "shab" {
		return shab_dto.EditMinNightBody{
			Action:  "set_min_days",
			Dates:   h.datesToJalali(dates),
			MinDays: amount,
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
	} else if h.service == "jajiga" {
		return jajiga_dto.MinNightBody{
			RoomID:    roomID,
			Dates:     dates,
			MinNights: 1,
		}
	} else if h.service == "otaghak" {
		return otaghak_dto.EditMinNightBody{
			MinNights:     1,
			EffectiveDays: h.datesToIso(dates),
			RoomID:        roomID,
		}
	} else if h.service == "shab" {
		return shab_dto.EditMinNightBody{
			Action:  "set_min_days",
			Dates:   h.datesToJalali(dates),
			MinDays: 1,
		}
	}
	return nil
}

func (h *homsaService) generateAuthResponse() interfaces.ApiResponseManager {
	if h.service == "homsa" {
		return &homsa_dto.HomsaAuthResponse{}
	} else if h.service == "jabama" {
		return &jabama_dto.Response{}
	} else if h.service == "jajiga" {
		return &jajiga_dto.AuthOkResponse{}
	} else if h.service == "otaghak" {
		return &otaghak_dto.AuthOkResponse{}
	} else if h.service == "shab" {
		return &shab_dto.AuthResponse{}
	}
	return nil
}

func (h *homsaService) generateOTPResponse() interfaces.ApiResponseManager {
	if h.service == "homsa" {
		return &homsa_dto.HomsaOTPResponse{}
	} else if h.service == "jabama" {
		return &jabama_dto.Response{}
	} else if h.service == "jajiga" {
		return &jajiga_dto.OTPResponse{}
	} else if h.service == "otaghak" {
		return &otaghak_dto.OTPResponse{}
	} else if h.service == "shab" {
		return &shab_dto.AuthOTPResponse{}
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
	} else if h.service == "jabama" {
		return &jabama_dto.Response{}
	} else if h.service == "jajiga" {
		return &jajiga_dto.ErrorResponse{}
	} else if h.service == "otaghak" {
		return &otaghak_dto.ErrorResponse{}
	} else if h.service == "shab" {
		return &shab_dto.ErrResponse{}
	}
	return nil
}
