package jabama_dto

type OTPLogin struct {
	Mobile string `json:"mobile"`
	Code   string `json:"code"`
}
