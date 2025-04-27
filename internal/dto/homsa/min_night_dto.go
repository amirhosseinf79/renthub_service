package homsa_dto

type HomsaSetMinNightBody struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Min       int    `json:"min"`
	Max       *int   `json:"max"`
}

type HomsaUnsetMinNightBody struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
