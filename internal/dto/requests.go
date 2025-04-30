package dto

type ReqHeaderEntry struct {
	UpdateId    string `json:"updateId"`
	CallbackUrl string `json:"callbackUrl"`
	ClientID    string `json:"userId"`
}

type SiteEntry struct {
	Site  string `json:"site"`
	Code  string `json:"code"`
	Price int    `json:"price"`
}
type DateEntry struct {
	Dates []string `json:"dates"`
}

// Request bodies
type EditPriceRequest struct {
	ReqHeaderEntry
	DateEntry
	Prices []SiteEntry `json:"prices"`
}

type EditDiscountRequest struct {
	ReqHeaderEntry
	DateEntry
	DiscountPercent int         `json:"discountPercent"`
	Sites           []SiteEntry `json:"sites"`
}

type EditMinNightRequest struct {
	ReqHeaderEntry
	DateEntry
	LimitDays int         `json:"limitDays"`
	Sites     []SiteEntry `json:"sites"`
}

type EditCalendarRequest struct {
	ReqHeaderEntry
	DateEntry
	Action string      `json:"action"`
	Sites  []SiteEntry `json:"sites"`
}
