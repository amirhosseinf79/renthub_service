package jabama_dto

type CalendarResponse struct {
	Result  Result `json:"result"`
	Success bool   `json:"success"`
}

type Result struct {
	Price    Price `json:"price"`
	Discount any   `json:"discount"` // can change to appropriate type if structure known
	Extra    any   `json:"extra"`    // can change to appropriate type if structure known
}

type Price struct {
	CurrentPrice    int         `json:"current_price"`
	SuggestedPrice  any         `json:"suggested_price"`  // nullable
	PricingGuidance any         `json:"pricing_guidance"` // nullable
	DefaultHint     DefaultHint `json:"default_hint"`
}

type DefaultHint struct {
	HasSimilarity bool   `json:"has_similarity"`
	Icon          string `json:"icon"`
	Text          string `json:"text"`
	Color         string `json:"color"`
}
