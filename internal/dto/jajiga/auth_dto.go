package jajiga_dto

type JajigaAuthRequestBody struct {
	Mobile     string  `json:"mobile"`
	Password   *string `json:"password"`
	ISO2       string  `json:"iso2"`
	ClientID   string  `json:"client_id"`
	ClientType string  `json:"client_type"`
}

type JajigaTokenAuthRequestBody struct {
	Mobile     string  `json:"mobile"`
	Token      *string `json:"token"`
	ISO2       string  `json:"iso2"`
	ClientID   string  `json:"client_id"`
	ClientType string  `json:"client_type"`
}

type OTPLogin struct {
	Mobile string `json:"mobile"`
	ISO2   string `json:"iso2"`
}
