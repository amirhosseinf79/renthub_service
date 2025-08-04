package homsa_dto

type homsaOTPData struct {
	New bool `json:"new"`
	TTL int  `json:"ttl"`
}

type HomsaLoginUserPass struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	UseOTP   bool   `json:"use_otp"`
}

type HomsaOTPLogin struct {
	Mobile string `json:"mobile"`
}
