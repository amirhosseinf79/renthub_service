package homsa_dto

type HomsaPriceBody struct {
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	Price        int    `json:"price"`
	KeepDiscount int    `json:"keep_discount"`
}
