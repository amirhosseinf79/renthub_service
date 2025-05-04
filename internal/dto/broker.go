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
