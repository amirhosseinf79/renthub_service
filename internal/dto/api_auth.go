package dto

type ApiAuthRequest struct {
	ClientID     string `json:"clientId" validate:"required"`
	Username     string `json:"username" validate:"required"`
	Service      string `json:"service" validate:"required,oneof=homsa mihmansho jabama shab jajiga otaghak"`
	AccessToken  string `json:"token" validate:"required"`
	Password     string `json:"password"`
	RefreshToken string `json:"refreshToken"`
	Ucode        string `json:"ucode"`
}

type OTPCreds struct {
	PhoneNumber string
	OTPCode     string
}
