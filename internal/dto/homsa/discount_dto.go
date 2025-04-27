package homsa_dto

type HomsaAddDiscountBody struct {
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	Discount     int    `json:"discount"`
	KeepDiscount int    `json:"keep_discount"`
}

type HomsaRemoveDiscountBody struct {
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	KeepDiscount int    `json:"keep_discount"`
}
