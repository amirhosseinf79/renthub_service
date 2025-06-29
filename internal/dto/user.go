package dto

type UserFilter struct {
	Email string `query:"email"`
}

type UserLogin struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
}

type UserRegister struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	HookToken   string `json:"hook_token" validate:"required"`
	HookRefresh string `json:"hook_refresh" validate:"required"`
	RefreshURL  string `json:"refresh_url" validate:"required"`
}

type UserUpdate struct {
	Email       *string `json:"email" validate:"required,email"`
	Password    *string `json:"password" validate:"omitempty"`
	FirstName   *string `json:"first_name" validate:"omitempty"`
	LastName    *string `json:"last_name" validate:"omitempty"`
	HookToken   *string `json:"hook_token" validate:"omitempty"`
	HookRefresh *string `json:"hook_refresh" validate:"omitempty"`
	RefreshURL  *string `json:"refresh_url" validate:"omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type OTPErrorResponse struct {
	Message        string `json:"message"`
	ServiceMessage string `json:"serviceMessage"`
}

type RefreshTokenBody struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
