package jajiga

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

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
	request        interfaces.FetchService
}

func New(
	apiAuthService interfaces.ApiAuthInterface,
	request interfaces.FetchService,
) interfaces.ApiService {
	return &service{
		service:        "jajiga",
		apiAuthService: apiAuthService,
		request:        request,
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
				"x-lang":          "fa",
				"Authorization":   "Bearer %v",
				"x-request-b":     "%v",
				"x-request-h":     "%v",
				"x-request-t":     "%v",
				"x-session-id":    "%v",
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

func (h *service) generateXHeaders(method, url string, body any) map[string]string {
	// method := "POST"
	// url := "/auth/otp/send"
	// body := `{"mobile":"09334429096","iso2":"IR"}`
	plusStr := "22ab046bd59212160e654bd5a610eb8da87723239db6d059f3074ad451c667b2"

	xRequestT := fmt.Sprintf("%d", time.Now().Unix())

	var b = []byte{}
	var err error

	if body != nil {
		b, err = json.Marshal(body)
		if err != nil {
			return nil
		}
	}

	md5Sum := md5.Sum(b)
	xRequestB := hex.EncodeToString(md5Sum[:])

	raw := strings.ToUpper(method) + "/api" + url + xRequestB + xRequestT + plusStr

	md5Raw := md5.Sum([]byte(raw))
	xRequestH := hex.EncodeToString(md5Raw[:])[:32]

	fmt.Println("x-request-t:", xRequestT)

	fmt.Println("body:", string(b))
	fmt.Println("x-request-b:", xRequestB)

	fmt.Println("row:", raw)
	fmt.Println("x-request-h:", xRequestH)

	return map[string]string{
		"x-request-t": xRequestT,
		"x-request-b": xRequestB,
		"x-request-h": xRequestH,
	}
}

func (h *service) generateAuthResponse() interfaces.ApiResponseManager {
	return &jajiga_dto.AuthOkResponse{}
}

func (h *service) generateOTPResponse() interfaces.ApiResponseManager {
	return &jajiga_dto.OTPResponse{}
}

func (h *service) generateUpdateErrResponse() interfaces.ApiResponseManager {
	return &jajiga_dto.ErrorResponse{}
}

func (h *service) generateErrResponse() interfaces.ApiResponseManager {
	return &jajiga_dto.ErrorResponse{}
}

// Body
func (h *service) generateCalendarBody(roomID string, setOpen bool, dates []string) jajiga_dto.CalendarBody {
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

func (h *service) generateAddDiscountBody(roomID string, amount int, dates []string) jajiga_dto.DiscountBody {
	return jajiga_dto.DiscountBody{
		RoomID:  roomID,
		Dates:   dates,
		Percent: amount,
	}
}

func (h *service) generateRemoveDiscountBody(roomID string, dates []string) jajiga_dto.DiscountBody {
	return jajiga_dto.DiscountBody{
		RoomID:  roomID,
		Dates:   dates,
		Percent: 0,
	}
}

func (h *service) generateEasyLoginBody(fields dto.ApiEasyLogin) jajiga_dto.JajigaAuthRequestBody {
	return jajiga_dto.JajigaAuthRequestBody{
		Mobile:     fields.Username,
		Password:   &fields.Password,
		ISO2:       "IR",
		ClientID:   uuid.New().String(),
		ClientType: "browser",
	}
}

func (h *service) generateSendOTPBody(phoneNumber string) jajiga_dto.OTPLogin {
	return jajiga_dto.OTPLogin{
		Mobile: phoneNumber,
		ISO2:   "IR",
	}
}

func (h *service) generateVerifyOTPBody(phoneNumber string, code string) jajiga_dto.JajigaTokenAuthRequestBody {
	return jajiga_dto.JajigaTokenAuthRequestBody{
		Mobile:   phoneNumber,
		Token:    &code,
		ClientID: uuid.New().String(),
		ISO2:     "IR",
	}
}

func (h *service) generateSetMinNightBody(roomID string, amount int, dates []string) jajiga_dto.MinNightBody {
	return jajiga_dto.MinNightBody{
		RoomID:    roomID,
		Dates:     dates,
		MinNights: amount,
	}
}

func (h *service) generateUnsetMinNightBody(roomID string, dates []string) jajiga_dto.MinNightBody {
	return jajiga_dto.MinNightBody{
		RoomID:    roomID,
		Dates:     dates,
		MinNights: 1,
	}
}

func (h *service) generatePriceBody(roomID string, amount int, dates []string) jajiga_dto.PriceBody {
	return jajiga_dto.PriceBody{
		RoomID: roomID,
		Dates:  dates,
		Price:  amount,
	}
}
