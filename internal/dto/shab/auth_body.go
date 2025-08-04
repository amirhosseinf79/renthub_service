package shab_dto

type OTPBody struct {
	Mobile      string `json:"mobile"`
	CountryCode string `json:"country_code"`
}

type VerifyOTOBody struct {
	Mobile      string `json:"mobile"`
	CountryCode string `json:"country_code"`
	Code        string `json:"code"`
}
