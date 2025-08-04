package otaghak_dto

type DayPricePair struct {
	Day   string `json:"day"`
	Price int    `json:"price"`
}

type CalendarBody struct {
	RoomID        string   `json:"roomId"`
	BlockedDays   []string `json:"blockedDays"`
	UnblockedDays []string `json:"unblockedDays"`
}

type EditPriceBody struct {
	RoomID       string         `json:"roomId"`
	PerDayPrices []DayPricePair `json:"perDayPrices"`
}

type EditDiscountBody struct {
	DiscountPercent       int      `json:"discountPercent"`
	EffectiveDays         []string `json:"effectiveDays"`
	EffectiveFromDateTime *string  `json:"effectiveFromDateTime"`
	IsActive              string   `json:"isActive"`
	RoomID                string   `json:"roomId"`
	EndDateTime           *string  `json:"endDateTime"`
	StartDateTime         *string  `json:"startDateTime"`
	EffectiveToDateTime   *string  `json:"effectiveToDateTime"`
}

type EditMinNightBody struct {
	EffectiveDays         []string `json:"effectiveDays"`
	EffectiveFromDateTime *string  `json:"effectiveFromDateTime"`
	IsActive              string   `json:"isActive"`
	RoomID                string   `json:"roomId"`
	MinNights             int      `json:"minNights"`
	EndDateTime           *string  `json:"endDateTime"`
	StartDateTime         *string  `json:"startDateTime"`
	EffectiveToDateTime   *string  `json:"effectiveToDateTime"`
}
