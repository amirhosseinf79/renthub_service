package mihmansho_dto

type MihmanshoErrorResponse struct {
	ErrorCode        int    `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
}

type MihmanshoProfileResponse struct {
	MihmanshoErrorResponse
	Cities   []any    `json:"Cities"`
	UserInfo userInfo `json:"UserInfo"`
}

type userInfo struct {
	Id                  int        `json:"Id"`
	Name                string     `json:"Name"`
	Family              string     `json:"Family"`
	FatherName          string     `json:"FatherName"`
	RegNumber           string     `json:"RegNumber"`
	CodeMeli            string     `json:"CodeMeli"`
	Gender              bool       `json:"Gender"`
	IsMarried           bool       `json:"IsMarried"`
	UserName            string     `json:"UserName"`
	Rate                int        `json:"Rate"`
	StateName           string     `json:"StateName"`
	StateId             int        `json:"StateId"`
	CityName            string     `json:"CityName"`
	CityId              int        `json:"CityId"`
	Email               string     `json:"Email"`
	Mobile              string     `json:"Mobile"`
	Mobile2             string     `json:"Mobile2"`
	Phone               string     `json:"Phone"`
	Address             string     `json:"Address"`
	AboutMe             string     `json:"AboutMe"`
	ZipCode             string     `json:"ZipCode"`
	VerifyAddress       bool       `json:"VerifyAddress"`
	VerifyCodeMeli      bool       `json:"VerifyCodeMeli"`
	VerifyEmail         bool       `json:"VerifyEmail"`
	VerifyMobile        bool       `json:"VerifyMobile"`
	Telegram            string     `json:"Telegram"`
	Instagram           string     `json:"Instagram"`
	CreationDate        string     `json:"CreationDate"`
	LastAccessDate      string     `json:"LastAccessDate"`
	Birthday            string     `json:"Birthday"`
	ImageProfile        string     `json:"ImageProfile"`
	ImageWall           string     `json:"ImageWall"`
	CreditAccount       int        `json:"CreditAccount"`
	UserType            int        `json:"UserType"`
	SuccessRequestCount int        `json:"SuccessRequestCount"`
	IsSuccessRequest    bool       `json:"IsSuccessRequest"`
	IsCancelRequest     bool       `json:"IsCancelRequest"`
	Credit              userCredit `json:"Credit"`
}

type userCredit struct {
	Label    string `json:"Label"`
	Value    string `json:"Value"`
	Currency string `json:"Currency"`
	TextLink string `json:"TextLink"`
}
