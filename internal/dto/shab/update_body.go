package shab_dto

type DayPricePair struct {
	Day   string `json:"day"`
	Price int    `json:"price"`
}

type CalendarBody struct {
	Action string   `json:"action"`
	Dates  []string `json:"dates"`
}

type EditPriceBody struct {
	Dates        []string `json:"dates"`
	Price        int      `json:"price"`
	KeepDiscount bool     `json:"keep_discount"`
}

type EditDiscountBody struct {
	Dates         []string `json:"dates"`
	Action        string   `json:"action"`
	DailyDiscount int      `json:"daily_discount"`
}

type UnsetDiscountBody struct {
	Dates  []string `json:"dates"`
	Action string   `json:"action"`
}

type EditMinNightBody struct {
	Dates   []string `json:"dates"`
	MinDays int      `json:"minimum_days"`
	Action  string   `json:"action"`
}
