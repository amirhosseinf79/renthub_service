package mihmansho_dto

type AuthBody struct {
	Username string `url:"Username"`
	Password string `url:"Password"`
}

type OTPBody struct {
	Mobile string `url:"mobile"`
	IsCode bool   `url:"isCode"`
}

type OTPVerifyBody struct {
	Mobile string `url:"mobile"`
	Code   string `url:"code"`
}
