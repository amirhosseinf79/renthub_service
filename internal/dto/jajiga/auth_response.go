package jajiga_dto

type AuthOkResponse struct {
	JWTToken string `json:"jwt_token"`
}

type OTPResponse struct {
	NotifyChannels []string `json:"notify_channels"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
