package request_v2

type ClientUpdateBody struct {
	UserID          uint
	Header          ReqHeaderEntry
	Services        []SiteEntry
	Dates           []string
	Action          string
	LimitDays       int
	DiscountPercent int
	WebhookBody     any
}

type ClientRecieveBody struct {
	UserID      uint
	Header      ReqHeaderEntry
	Services    []SiteRecieve
	WebhookBody any
}

type OTPBody struct {
	UserID      uint
	ClientID    string
	Service     string
	PhoneNumebr string
	Code        string
}
