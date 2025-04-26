package homsa

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/repository"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

type homsaService struct {
	apiAuthRepo repository.ApiAuthRepository
	service     string
	apiUrl      string
	endpoints   *dto.ApiEndpoints
}

func NewHomsaService(apiAuthRepo repository.ApiAuthRepository) interfaces.ApiService {
	return &homsaService{
		apiAuthRepo: apiAuthRepo,
		service:     "homsa",
		apiUrl:      "https://www.homsa.net/api/v2",
		endpoints: &dto.ApiEndpoints{
			LoginFirstStep:  "/newAuth/otp/send",
			LoginSecondStep: "/newAuth/login",
			LoginWithPass:   "/newAuth/login",
			GetProfile:      "/host/profile",
			OpenCalendar:    "/host/room/%v/calendar/unblock",
			CloseCalendar:   "/host/room/%v/calendar/block",
			EditPricePerDay: "/host/room/%v/calendar/set_price",
			AddDiscount:     "/host/room/%v/calendar/set_discount",
			RemoveDiscount:  "/host/room/%v/calendar/remove_discount",
			SetMinNight:     "/host/room/%v/calendar/update_availability_rule",
			UnsetMinNight:   "/host/room/%v/calendar/delete_availability_rule",
		},
	}
}

func (h *homsaService) GetHeader() map[string]string {
	return map[string]string{
		"user-Agent":      "Dart/2.19 (dart:io)",
		"accept":          "application/json",
		"accept-Encoding": "gzip",
		"host":            "www.homsa.net",
		"content-type":    "application/json; charset=UTF-8",
		"accept-charset":  "UTF-8",
		"authorization":   "bearer %v",
		"lang":            "fa",
	}
}
