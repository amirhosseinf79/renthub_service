package mihmansho_dto

type CalendarDetailsResponse struct {
	MihmanshoErrorResponse
	CalendarData  []CalendarDay  `json:"CalendarData"`
	AvailableDays []AvailableDay `json:"AvailableDays"`
	ReservedDays  []any          `json:"ReservedDays"`
	HolyDays      []any          `json:"HolyDays"`
}

type CalendarDay struct {
	Day          int     `json:"Day"`
	Price        string  `json:"Price"`
	NewPrice     *string `json:"NewPrice"`
	AddedPrice   string  `json:"AddedPrice"`
	Available    bool    `json:"Available"`
	UnAvailable  bool    `json:"UnAvailable"`
	HolyDay      bool    `json:"HolyDay"`
	Reserved     bool    `json:"Reserved"`
	TypeReserved int     `json:"TypeReserved"`
	ReadyReserve bool    `json:"ReadyReserve"`
	DiscountHost int     `json:"DiscountHost"`
}

type AvailableDay struct {
	ID                     int     `json:"Id"`
	Day                    int     `json:"Day"`
	Price                  string  `json:"Price"`
	PriceOld               *string `json:"PriceOld"`
	ExteraPrice            *string `json:"ExteraPrice"`
	PriceHost              int     `json:"PriceHost"`
	TypeWage               int     `json:"TypeWage"`
	Wage                   int     `json:"Wage"`
	TaxHost                int     `json:"TaxHost"`
	AddedPrice             string  `json:"AddedPrice"`
	AddedPriceOld          *string `json:"AddedPriceOld"`
	AddedExteraPrice       *string `json:"AddedExteraPrice"`
	Percent                int     `json:"Percent"`
	DiscountHost           int     `json:"DiscountHost"`
	Description            *string `json:"Description"`
	IsLastSecond           bool    `json:"IsLastSecond"`
	Date                   string  `json:"Date"`
	RequestId              int     `json:"RequestId"`
	IsReserveReady         bool    `json:"IsReserveReady"`
	IsWarrantyReadyReserve bool    `json:"IsWarrantyReadyReserve"`
	CancelRequest          int     `json:"CancelRequest"`
	MinDay                 int     `json:"MinDay"`
	MaxDay                 int     `json:"MaxDay"`
	IsEdited               bool    `json:"IsEdited"`
}
