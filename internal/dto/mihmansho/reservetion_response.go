package mihmansho_dto

type ReservationResponse struct {
	MihmanshoErrorResponse
	Type                    int                   `json:"Type"`
	TypeProduct             int                   `json:"TypeProduct"`
	TypeProductName         string                `json:"TypeProductName"`
	PageNumber              int                   `json:"PageNumber"`
	PageSize                int                   `json:"PageSize"`
	HasPageNext             bool                  `json:"hasPageNext"`
	Filters                 []ReservationFilter   `json:"Filters"`
	CancelRequestItems      []CancelRequestItem   `json:"CancelRequestItems"`
	CancelTermsAndCondition int                   `json:"CancelTermsAndCondition"`
	ListRequest             []ReservationListItem `json:"ListRequest"`
}

func (r *ReservationResponse) GetList() []ReservationListItem {
	return r.ListRequest
}

type ReservationFilter struct {
	Status int    `json:"Status"`
	Value  string `json:"Value"`
	Badge  string `json:"Badge"`
}

type CancelRequestItem struct {
	Key   int    `json:"Key"`
	Value string `json:"Value"`
}

type ReservationListItem struct {
	Id                       int     `json:"Id"`
	FollowUpNum              string  `json:"FollowUpNum"`
	ProductId                int     `json:"ProductId"`
	ProductName              string  `json:"ProductName"`
	Mobile                   string  `json:"Mobile"`
	NameOfUser               string  `json:"NameOfUser"`
	UserId                   int     `json:"UserId"`
	ImageProfile             string  `json:"ImageProfile"`
	OrderDate                string  `json:"OrderDate"`
	TypeName                 string  `json:"TypeName"`
	TypeValue                string  `json:"TypeValue"`
	Price                    string  `json:"Price"`
	TextPrice                string  `json:"TextPrice"`
	CountGuest               int     `json:"CountGuest"`
	StatusRequest            string  `json:"StatusRequest"`
	TypeRequest              int     `json:"TypeRequest"`
	TypeRequestName          string  `json:"TypeRequestName"`
	StatusPayment            int     `json:"StatusPayment"`
	StatusPaymentName        string  `json:"StatusPaymentName"`
	Status                   int     `json:"Status"`
	StatusName               string  `json:"StatusName"`
	TimeExpired              *string `json:"TimeExpired"`
	VerifyReciver            int     `json:"VerifyReciver"`
	VerifyReciverDate        string  `json:"VerifyReciverDate"`
	VerifySender             int     `json:"VerifySender"`
	Description              string  `json:"Description"`
	UrlImage                 string  `json:"UrlImage"`
	DateHolding              string  `json:"DateHolding"`
	EnterDate                string  `json:"EnterDate"`
	Duration                 string  `json:"Duration"`
	Location                 string  `json:"Location"`
	CountGuestText           string  `json:"CountGuestText"`
	TypeProduct              int     `json:"TypeProduct"`
	UrlDetail                string  `json:"UrlDetail"`
	UrlPayment               string  `json:"UrlPayment"`
	UrlComment               string  `json:"UrlComment"`
	ShowChat                 bool    `json:"ShowChat"`
	ShowSendImage            bool    `json:"ShowSendImage"`
	UrlSendImage             string  `json:"UrlSendImage"`
	ShowCancelRequest        bool    `json:"ShowCancelRequest"`
	StatusCancelRequest      int     `json:"StatusCancelRequest"`
	ButtonTextCancelRequest  string  `json:"ButtonTextCancelRequest"`
	DisableCancelRequest     bool    `json:"DisableCancelRequest"`
	DescriptionCancelRequest string  `json:"DescriptionCancelRequest"`
	UrlStatus                string  `json:"UrlStatus"`
	UrlStatusType            int     `json:"UrlStatusType"`
	Renew                    string  `json:"Renew"`
	Step1                    string  `json:"Step1"`
	Step2                    string  `json:"Step2"`
	Step3                    string  `json:"Step3"`
	Step4                    string  `json:"Step4"`
	UrlCancelRequest         string  `json:"UrlCancelRequest"`
}
