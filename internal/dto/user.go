package dto

type UserFilter struct {
	Email string `query:"email"`
}

type UserLogin struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
}

type UserRegister struct {
	Email     string `form:"email" validate:"required,email"`
	Password  string `form:"password" validate:"required"`
	FirstName string `form:"first_name" validate:"required"`
	LastName  string `form:"last_name" validate:"required"`
}

type ErrorResponse struct {
	Message        string
	ServiceMessage string
}
