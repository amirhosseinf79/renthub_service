package mihmansho

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	mihmansho_dto "github.com/amirhosseinf79/renthub_service/internal/dto/mihmansho"
	"github.com/amirhosseinf79/renthub_service/pkg"
)

type service struct {
	apiAuthService interfaces.ApiAuthInterface
	service        string
	apiSettings    dto.ApiSettings
	request        interfaces.FetchService
	chromium       interfaces.ChromeService
}

func New(
	apiAuthService interfaces.ApiAuthInterface,
	request interfaces.FetchService,

	chrome interfaces.ChromeService,
) interfaces.ApiService {
	return &service{
		service:        "mihmansho",
		apiAuthService: apiAuthService,
		request:        request,
		chromium:       chrome,
		apiSettings: dto.ApiSettings{
			ApiURL: "https://www.mihmansho.com",
			Endpoints: dto.ApiEndpoints{
				LoginFirstStep:  dto.EndP{Address: "/myapi/v1/loginfirststep", Method: "POST", ContentType: "query"},
				LoginSecondStep: dto.EndP{Address: "/myapi/v1/loginwithcode", Method: "POST", ContentType: "query"},
				LoginWithPass:   dto.EndP{Address: "/myapi/v1/login", Method: "POST", ContentType: "query"},
				GetProfile:      dto.EndP{Address: "/myapi/v1/getprofile", Method: "GET", ContentType: "body"},
				OpenCalendar:    dto.EndP{Address: "/myapi/v1/ReserveDates", Method: "POST", ContentType: "multipart"},
				CloseCalendar:   dto.EndP{Address: "/myapi/v1/ReserveDates", Method: "POST", ContentType: "multipart"},
				EditPricePerDay: dto.EndP{Address: "/myapi/v1/EditPricesCalendar?ProductId=%v&Price=%v&AddedGuestPrice=0", Method: "POST", ContentType: "multipart"},
				AddDiscount:     dto.EndP{Address: "/Account/Home/HostDiscountSave", Method: "POST", ContentType: "multipart"},
				RemoveDiscount:  dto.EndP{Address: "/Account/Home/HostDiscountSave", Method: "POST", ContentType: "multipart"},
				SetMinNight:     dto.EndP{Address: "/Account/Home/AddSpecificsMinDay", Method: "POST", ContentType: "multipart"},
				UnsetMinNight:   dto.EndP{Address: "/Account/Home/AddSpecificsMinDay", Method: "POST", ContentType: "multipart"},
			},
			Headers: map[string]string{
				"user-agent":      "okhttp/3.12.1",
				"accept-encoding": "gzip",
				"version":         "1.3.6",
				"city":            "0",
				"ucode":           "%v",
				"token":           "%v",
				"imei":            "%v",
				"cookie":          "ASP.NET_SessionId=%v",
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
		"ucode":  token.Ucode,
		"token":  token.AccessToken,
		"imei":   "c00add1deb77e991",
		"cookie": token.RefreshToken,
	}
}

func (h *service) generateAuthResponse() interfaces.ApiResponseManager {
	return &mihmansho_dto.AuthResponse{}
}

func (h *service) generateOTPResponse() interfaces.ApiResponseManager {
	return &mihmansho_dto.MihmanshoErrorResponse{}
}

// func (h *service) generateProfileResponse() interfaces.ApiResponseManager {
// 	return &mihmansho_dto.MihmanshoProfileResponse{}
// }

func (h *service) generateUpdateErrResponse() interfaces.ApiResponseManager {
	return &mihmansho_dto.MihmanshoErrorResponse{}
}

func (h *service) generateErrResponse() interfaces.ApiResponseManager {
	return &mihmansho_dto.MihmanshoErrorResponse{}
}

// Body
func (h *service) generateCalendarBody(roomID string, setOpen bool, dates []string) []byte {
	fDate := mihmansho_dto.Calendar{
		ProductId: roomID,
	}
	fmt.Println("mihmansho", dates)
	jdates := pkg.DatesToJalali(dates, false)
	for _, jdate := range jdates {
		fDate.Dates = append(fDate.Dates, mihmansho_dto.CalendarDates{Date: jdate, IsReserve: !setOpen, RequestId: 0})
	}
	bdata, err := json.Marshal(fDate)
	if err != nil {
		return nil
	}
	mainBody := mihmansho_dto.FormBody{
		"Dates": string(bdata),
	}
	mbody, err := json.Marshal(mainBody)
	if err != nil {
		return nil
	}
	return mbody
}

// func (h *service) generateEasyLoginBody(fields dto.ApiEasyLogin) mihmansho_dto.AuthBody {
// 	return mihmansho_dto.AuthBody{
// 		Username: fields.Username,
// 		Password: fields.Password,
// 	}
// }

func (h *service) generateSendOTPBody(phoneNumber string) mihmansho_dto.OTPBody {
	return mihmansho_dto.OTPBody{
		Mobile: phoneNumber,
		IsCode: true,
	}
}

func (h *service) generateVerifyOTPBody(phoneNumber string, code string) mihmansho_dto.OTPVerifyBody {
	return mihmansho_dto.OTPVerifyBody{
		Mobile: phoneNumber,
		Code:   code,
	}
}

func (h *service) generatePriceBody(dates []string) []byte {
	pbody := mihmansho_dto.FormBody{}
	jdates := pkg.DatesToJalali(dates, false)
	for i, date := range jdates {
		pbody[fmt.Sprintf("Dates_num%v", i)] = date
	}
	mbody, err := json.Marshal(pbody)
	if err != nil {
		return nil
	}
	return mbody
}

func (h *service) generateMinNightBody(roomID string, dates []string, minDay int) []byte {

	pbody := mihmansho_dto.FormBody{}
	jdates := pkg.DatesToJalali(dates, false)
	for i, date := range jdates {
		pbody[fmt.Sprintf("Date_num%v", i)] = date
	}
	pbody["ProductId"] = roomID
	pbody["MinDay"] = fmt.Sprintf("%v", minDay)
	mbody, err := json.Marshal(pbody)
	if err != nil {
		return nil
	}
	return mbody
}

func (h *service) generateDiscountBody(roomID string, dates []string, amount int) []byte {
	var active string
	if amount > 0 {
		active = "true"
	} else {
		active = "false"
	}
	jdates := pkg.DatesToJalali(dates, false)
	pbody := mihmansho_dto.FormBody{}
	pbody["dh.ProductId"] = roomID
	// pbody["dh.ActiveDateDiscountHost"] = active
	// pbody["dh.StringStartDateDiscountHost"] = jdates[0]
	// pbody["dh.StringEndDateDiscountHost"] = jdates[len(jdates)-1]
	// pbody["dh.PercentDateDiscountHost"] = fmt.Sprintf("%v", amount)

	for i, date := range jdates {
		pbody[fmt.Sprintf("dhh[%v].Active", i)] = active
		pbody[fmt.Sprintf("dhh[%v].StringStartDate", i)] = date
		pbody[fmt.Sprintf("dhh[%v].StringEndDate", i)] = date
	}

	mbody, err := json.Marshal(pbody)
	if err != nil {
		return nil
	}
	return mbody
}
