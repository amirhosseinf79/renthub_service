package dto

type ClientUpdateBody struct {
	UserID          uint
	Header          ReqHeaderEntry
	Services        []SiteEntry
	Dates           []string
	Action          string
	LimitDays       int
	DiscountPercent int
	FinalResult     ManagerResponse
}

type OTPBody struct {
	UserID      uint
	ClientID    string
	Service     string
	PhoneNumebr string
	Code        string
}
