package otaghak_dto

type AuthOkResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type OTPResponse struct {
	Odata           string `json:"@odata.context"`
	TimeExpiration  int    `json:"timeExpiration"`
	CellPhoneNumber string `json:"cellPhoneNumber"`
}

type ErrorResponse struct {
	Message          string `json:"message"`
	Code             string `json:"code"`
	TechnicalMessage string `json:"technical_message"`
	HttpResponseCode int    `json:"http_response_code"`
	CorrelationId    string `json:"correlation_id"`
	Detail           string `json:"detail"`
}
