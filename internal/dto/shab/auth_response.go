package shab_dto

type authOk struct {
	AccessToken string `json:"access_token"`
}

type AuthOTPResponse struct {
	Meta meta `json:"meta"`
}

type AuthResponse struct {
	Data authOk `json:"data"`
	Meta meta   `json:"meta"`
}

type ErrResponse struct {
	Meta meta `json:"meta"`
}

type meta struct {
	Status int `json:"status"`
}
