package mihmansho

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/domain/repository"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	mihmansho_dto "github.com/amirhosseinf79/renthub_service/internal/dto/mihmansho"
	"github.com/amirhosseinf79/renthub_service/pkg"
)

type service struct {
	apiAuthRepo repository.ApiAuthRepository
	service     string
	apiSettings dto.ApiSettings
}

func New(apiAuthRepo repository.ApiAuthRepository) interfaces.ApiService {
	return &service{
		service:     "mihmansho",
		apiAuthRepo: apiAuthRepo,
		apiSettings: dto.ApiSettings{
			ApiURL: "https://www.mihmansho.com/myapi/v1",
			Endpoints: dto.ApiEndpoints{
				LoginFirstStep:  dto.EndP{Address: "/loginfirststep", Method: "POST", ContentType: "query"},
				LoginSecondStep: dto.EndP{Address: "/loginwithcode", Method: "POST", ContentType: "query"},
				LoginWithPass:   dto.EndP{Address: "/login", Method: "POST", ContentType: "query"},
				GetProfile:      dto.EndP{Address: "/getprofile", Method: "GET", ContentType: "body"},
				OpenCalendar:    dto.EndP{Address: "/ReserveDates", Method: "POST", ContentType: "multipart"},
				CloseCalendar:   dto.EndP{Address: "/ReserveDates", Method: "POST", ContentType: "multipart"},
				EditPricePerDay: dto.EndP{Address: "/EditPricesCalendar?ProductId=%v&Price=%v&AddedGuestPrice=0", Method: "POST", ContentType: "multipart"},
			},
			Headers: map[string]string{
				"user-agent":      "okhttp/3.12.1",
				"accept-encoding": "gzip",
				"version":         "1.3.6",
				"city":            "0",
				"ucode":           "%v",
				"token":           "%v",
				"imei":            "%v",
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
		"ucode": token.Ucode,
		"token": token.AccessToken,
		"imei":  "c00add1deb77e991",
	}
}

func (h *service) generateAuthResponse() interfaces.ApiResponseManager {
	return &mihmansho_dto.AuthResponse{}
}

func (h *service) generateOTPResponse() interfaces.ApiResponseManager {
	return &mihmansho_dto.MihmanshoErrorResponse{}
}

func (h *service) generateProfileResponse() interfaces.ApiResponseManager {
	return &mihmansho_dto.MihmanshoProfileResponse{}
}

func (h *service) generateUpdateErrResponse() interfaces.ApiResponseManager {
	return &mihmansho_dto.MihmanshoErrorResponse{}
}

func (h *service) generateErrResponse() interfaces.ApiResponseManager {
	return &mihmansho_dto.MihmanshoErrorResponse{}
}

// Body
func (h *service) generateCalendarBody(roomID string, setOpen bool, dates []string) any {
	fDate := mihmansho_dto.Calendar{
		ProductId: roomID,
	}
	jdates := pkg.DatesToJalali(dates, false)
	for _, jdate := range jdates {
		fDate.Dates = append(fDate.Dates, mihmansho_dto.CalendarDates{Date: jdate, IsReserve: !setOpen, RequestId: 0})
	}
	bdata, err := json.Marshal(fDate)
	if err != nil {
		return err
	}
	mainBody := mihmansho_dto.FormBody{
		"Dates": string(bdata),
	}
	mbody, err := json.Marshal(mainBody)
	if err != nil {
		return err
	}
	return mbody
}

func (h *service) generateAddDiscountBody() any {
	return nil
}

func (h *service) generateRemoveDiscountBody() any {
	return nil
}

func (h *service) generateEasyLoginBody(fields dto.ApiEasyLogin) any {
	return mihmansho_dto.AuthBody{
		Username: fields.Username,
		Password: fields.Password,
	}
}

func (h *service) generateSendOTPBody(phoneNumber string) any {
	return mihmansho_dto.OTPBody{
		Mobile: phoneNumber,
		IsCode: true,
	}
}

func (h *service) generateVerifyOTPBody(phoneNumber string, code string) any {
	return mihmansho_dto.OTPVerifyBody{
		Mobile: phoneNumber,
		Code:   code,
	}
}

func (h *service) generateSetMinNightBody() any {
	return nil
}

func (h *service) generateUnsetMinNightBody() any {
	return nil
}

func (h *service) generatePriceBody(dates []string) any {
	pbody := mihmansho_dto.FormBody{}
	jdates := pkg.DatesToJalali(dates, false)
	for _, date := range jdates {
		pbody["Dates"] = date
	}
	mbody, err := json.Marshal(pbody)
	if err != nil {
		return err
	}
	return mbody
}
