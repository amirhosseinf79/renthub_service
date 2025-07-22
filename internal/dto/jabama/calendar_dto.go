package jabama_dto

type OpenClosCalendar struct {
	Dates []string `json:"dates"`
}

type EditPricePerDay struct {
	Type  *string  `json:"type"`
	Days  []string `json:"days"`
	Value int      `json:"value"`
}

type DiscountDTO struct {
	Days  []string `json:"days"`
	Value int      `json:"value"`
}
