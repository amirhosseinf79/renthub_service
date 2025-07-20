package otaghak

import (
	"bytes"
	"fmt"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	otaghak_dto "github.com/amirhosseinf79/renthub_service/internal/dto/otaghak"
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
		service:        "otaghak",
		apiAuthService: apiAuthService,
		request:        request,
		apiSettings: dto.ApiSettings{
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
				"Authorization":   "Bearer %v",
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
	return &otaghak_dto.AuthOkResponse{}
}

func (h *service) generateOTPResponse() interfaces.ApiResponseManager {
	return &otaghak_dto.OTPResponse{}
}

func (h *service) generateUpdateErrResponse() interfaces.ApiResponseManager {
	return &otaghak_dto.ErrorResponse{}
}

func (h *service) generateErrResponse() interfaces.ApiResponseManager {
	return &otaghak_dto.ErrorResponse{}
}

// Body
func (h *service) generateCalendarBody(roomID string, setOpen bool, dates []string) otaghak_dto.CalendarBody {
	if setOpen {
		return otaghak_dto.CalendarBody{
			RoomID:        roomID,
			UnblockedDays: pkg.DatesToIso(dates),
		}
	}
	return otaghak_dto.CalendarBody{
		RoomID:      roomID,
		BlockedDays: pkg.DatesToIso(dates),
	}
}

func (h *service) generateAddDiscountBody(roomID string, amount int, dates []string) otaghak_dto.EditDiscountBody {
	return otaghak_dto.EditDiscountBody{
		DiscountPercent: amount,
		EffectiveDays:   pkg.DatesToIso(dates),
		RoomID:          roomID,
	}
}

func (h *service) generateRemoveDiscountBody(roomID string, dates []string) otaghak_dto.EditDiscountBody {
	return otaghak_dto.EditDiscountBody{
		DiscountPercent: 0,
		EffectiveDays:   pkg.DatesToIso(dates),
		RoomID:          roomID,
	}
}

func (h *service) generateEasyLoginBody(fields dto.ApiEasyLogin) otaghak_dto.OtaghakAuthRequestBody {
	return otaghak_dto.OtaghakAuthRequestBody{
		UserName:     fields.Username,
		Password:     fields.Password,
		ClientId:     "Otaghak",
		ClientSecret: "secret",
		ArcValues:    map[string]string{},
	}
}

func (h *service) generateSendOTPBody(phoneNumber string) otaghak_dto.OTPBody {
	return otaghak_dto.OTPBody{
		UserName:   phoneNumber,
		IsShortOtp: true,
	}
}

func (h *service) generateVerifyOTPBody(phoneNumber string, code string) otaghak_dto.OtaghakAuthRequestBody {
	return otaghak_dto.OtaghakAuthRequestBody{
		UserName:     phoneNumber,
		ClientId:     "Otaghak",
		ClientSecret: "secret",
		ArcValues:    map[string]string{"OtpCode": code},
	}
}

func (h *service) generateSetMinNightBody(roomID string, amount int, dates []string) otaghak_dto.EditMinNightBody {
	return otaghak_dto.EditMinNightBody{
		MinNights:     amount,
		EffectiveDays: pkg.DatesToIso(dates),
		RoomID:        roomID,
	}
}

func (h *service) generateUnsetMinNightBody(roomID string, dates []string) otaghak_dto.EditMinNightBody {
	return otaghak_dto.EditMinNightBody{
		MinNights:     1,
		EffectiveDays: pkg.DatesToIso(dates),
		RoomID:        roomID,
	}
}

func (h *service) generatePriceBody(roomID string, amount int, dates []string) otaghak_dto.EditPriceBody {
	formattedDates := pkg.DatesToIso(dates)
	var formattedDays []otaghak_dto.DayPricePair
	for _, item := range formattedDates {
		formattedDays = append(formattedDays, otaghak_dto.DayPricePair{Day: item, Price: amount})
	}
	return otaghak_dto.EditPriceBody{
		RoomID:       roomID,
		PerDayPrices: formattedDays,
	}
}
