package jajiga_dto

type CalendarBody struct {
	RoomID       string   `json:"room_id"`
	Dates        []string `json:"dates"`
	DisableCount int      `json:"disable_count"`
}

type PriceBody struct {
	RoomID string   `json:"room_id"`
	Dates  []string `json:"dates"`
	Price  int      `json:"price"`
}

type DiscountBody struct {
	RoomID  string   `json:"room_id"`
	Dates   []string `json:"dates"`
	Percent int      `json:"discount_percent"`
}

type MinNightBody struct {
	RoomID    string   `json:"room_id"`
	Dates     []string `json:"dates"`
	MinNights int      `json:"stays_min"`
}
