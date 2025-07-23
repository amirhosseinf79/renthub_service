package mihmansho_dto

type RoomDetailsResponse struct {
	ErrorCode        int     `json:"errorCode"`
	ErrorDescription string  `json:"errorDescription"`
	Details          Details `json:"Details"`
}

type Details struct {
	PrePreview        bool   `json:"PrePreview"`
	PrePreviewMessage string `json:"PrePreviewMessage"`
	Id                int    `json:"Id"`
	Name              string `json:"Name"`
	Rate              string `json:"Rate"`
	CityName          string `json:"CityName"`
	StateName         string `json:"StateName"`
	Region            string `json:"Region"`
	Price             int    `json:"Price"`
	PaymentNote       string `json:"PaymentNote"`
}
