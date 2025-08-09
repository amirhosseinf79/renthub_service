package request_v2

type ReqHeaderEntry struct {
	UpdateId    string `json:"updateId" validate:"required"`
	CallbackUrl string `json:"callbackUrl" validate:"required"`
	ClientID    string `json:"clientId,omitempty"`
}

type ReqHeaderWithClientEntry struct {
	ReqHeaderEntry
	ClientID string `json:"clientId" validate:"required"`
}

type SiteRecieve struct {
	ClientID string   `json:"clientId" validate:"required"`
	Site     string   `json:"site" validate:"required"`
	Filters  []string `json:"filters"`
}

type SiteEntry struct {
	ClientID string `json:"clientId" validate:"required"`
	Site     string `json:"site" validate:"required"`
	Code     string `json:"code" validate:"required"`
	Price    int    `json:"price"`
}

type DateEntry struct {
	Dates []string `json:"dates" validate:"required"`
}

// Request bodies
type EditPriceRequest struct {
	ReqHeaderEntry
	DateEntry
	Prices []SiteEntry `json:"prices" validate:"required"`
}

type EditDiscountRequest struct {
	ReqHeaderEntry
	DateEntry
	DiscountPercent int         `json:"discountPercent" validate:"gte=0"`
	Sites           []SiteEntry `json:"sites" validate:"required"`
}

type EditMinNightRequest struct {
	ReqHeaderEntry
	DateEntry
	LimitDays int         `json:"limitDays" validate:"gte=1"`
	Sites     []SiteEntry `json:"sites" validate:"required"`
}

type EditCalendarRequest struct {
	ReqHeaderEntry
	DateEntry
	Action string      `json:"action" validate:"oneof=block unblock"`
	Sites  []SiteEntry `json:"sites" validate:"required"`
}

type RefreshTokenRequest struct {
	ReqHeaderEntry
	Sites []SiteEntry `json:"sites" validate:"required"`
}

type OTPSendRequest struct {
	ClientID    string `json:"clientId" validate:"required"`
	Service     string `json:"service" validate:"required"`
	PhoneNumebr string `json:"phoneNumber" validate:"required"`
}

type OTPVerifyRequest struct {
	ClientID    string `json:"clientId" validate:"required"`
	Service     string `json:"service" validate:"required"`
	PhoneNumebr string `json:"phoneNumber" validate:"required"`
	Code        string `json:"code" validate:"required"`
}
