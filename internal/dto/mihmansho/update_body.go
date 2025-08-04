package mihmansho_dto

type CalendarDates struct {
	Date      string `json:"Date"`
	IsReserve bool   `json:"IsReserve"`
	RequestId int    `json:"RequestId"`
}

type Calendar struct {
	Dates     []CalendarDates `json:"Dates"`
	ProductId string          `json:"ProductId"`
}

type FormBody map[string]string
