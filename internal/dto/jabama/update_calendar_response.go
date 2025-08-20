package jabama_dto

import "time"

type UpdateCalendarStatusResponse struct {
	Result UpdateCalendarResult `json:"result"`
}

type UpdateCalendarResult struct {
	ID    int                 `json:"id"`
	Price UpdateCalendarPrice `json:"price"`
}

type UpdateCalendarPrice struct {
	Custom []CustomUpdateCalendarPrice `json:"custom"`
}

type CustomUpdateCalendarPrice struct {
	Status         string    `json:"status"`
	Price          float64   `json:"price"`
	Discount       float64   `json:"discount"`
	JabamaDiscount float64   `json:"jabamaDiscount"`
	Date           time.Time `json:"date"`
	ExtraPeople    float64   `json:"extraPeople"`
}
