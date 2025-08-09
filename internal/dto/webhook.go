package dto

type WebhookFields struct {
	UserID      uint
	UpdateId    string
	CallbackUrl string
	Body        any
}

type WebhookRefreshBody struct {
	RefreshToken string `json:"refreshToken"`
}

type WebhookRefreshResponse struct {
	AccessToken string `json:"accessToken"`
}
