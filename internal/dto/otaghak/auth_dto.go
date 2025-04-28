package otaghak_dto

type OtaghakAuthRequestBody struct {
	UserName     string            `json:"userName"`
	Password     string            `json:"password"`
	ClientId     string            `json:"clientId"`
	ClientSecret string            `json:"clientSecret"`
	ArcValues    map[string]string `json:"arcValues"`
}

type OTPBody struct {
	UserName   string `json:"username"`
	IsShortOtp bool   `json:"isShortOtp"`
}

type VeriftOTOBody struct {
	UserName string `json:"userName"`
	Code     string `json:"code"`
}
