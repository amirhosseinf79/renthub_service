package dto

type WebhookRefreshBody struct {
	RefreshToken string `json:"refreshToken"`
}

type WebhookRefreshResponse struct {
	AccessToken string `json:"accessToken"`
}
