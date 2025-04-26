package dto

type HomsaLoginUserPass struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	UseOTP   bool   `json:"use_otp"`
}

type HomsaAuthResponse struct {
	UserID       int    `json:"user_id"`
	Email        string `json:"email"`
	FullName     string `json:"full_name"`
	ImageURL     string `json:"image_url"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpireAt     int    `json:"expire_at"`
	PhoneNumber  string `json:"phone_number"`
}

type HomsaErrorResponse struct {
	Code    string              `json:"code"`
	Message string              `json:"message"`
	Errors  map[string][]string `json:"errors"`
}
