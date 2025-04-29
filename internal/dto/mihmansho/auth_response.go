package mihmansho_dto

type AuthResponse struct {
	UCode            string `json:"ucode"`
	Token            string `json:"Token"`
	UserType         int    `json:"UserType"`
	ErrorCode        int    `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
}
