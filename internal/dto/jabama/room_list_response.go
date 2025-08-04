package jabama_dto

import "github.com/amirhosseinf79/renthub_service/internal/domain/models"

type RoomListResponse struct {
	Result struct {
		Items []struct {
			ID                   string `json:"id"`
			Code                 int    `json:"code"`
			Title                string `json:"title"`
			StatusText           string `json:"statusText"`
			Status               string `json:"status"`
			PlaceType            string `json:"placeType"`
			ComplexID            string `json:"complexId"`
			LastUpdateText       string `json:"lastUpdateText"`
			NeedToBeUpdated      bool   `json:"needToBeUpdated"`
			IsSellable           bool   `json:"isSellable"`
			IsComplex            bool   `json:"isComplex"`
			UnitCount            int    `json:"unitCount"`
			AffiliateLink        string `json:"affiliateLink"`
			AffiliateDescription string `json:"affiliateDescription"`
			SmartPricing         bool   `json:"smartPricing"`
			ReservationType      string `json:"reservationType"`
		} `json:"items"`
	} `json:"result"`
	TargetURL           *string    `json:"targetUrl"`
	Success             bool       `json:"success"`
	Error               *authError `json:"error"`
	UnauthorizedRequest bool       `json:"unauthorizedRequest"`
	Wrapped             bool       `json:"__wrapped"`
	TraceID             string     `json:"__traceId"`
}

func (h *RoomListResponse) GetResult() (ok bool, result string) {
	ok = h.Error == nil
	result = "success"
	if !ok {
		result = h.Error.Message
	}
	return ok, result
}

func (h *RoomListResponse) GetToken() *models.ApiAuth {
	return &models.ApiAuth{}
}
