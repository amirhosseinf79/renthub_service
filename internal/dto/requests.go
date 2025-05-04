package dto

type ReqHeaderEntry struct {
	UpdateId    string `json:"updateId" validate:"required"`
	CallbackUrl string `json:"callbackUrl" validate:"required"`
	ClientID    string `json:"clientId" validate:"required"`
}

type SiteEntry struct {
	Site  string `json:"site" validate:"required"`
	Code  string `json:"code" validate:"required"`
	Price int    `json:"price"`
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
