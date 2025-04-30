package dto

type ApiAuthRequest struct {
	ClientID     string `json:"clientId" validate:"required"`
	Username     string `json:"username" validate:"required"`
	Service      string `json:"service" validate:"required"`
	AccessToken  string `json:"token" validate:"required"`
	Password     string `json:"password"`
	RefreshToken string `json:"refreshToken"`
	Ucode        string `json:"ucode"`
}
