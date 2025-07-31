package dto

type LogFilters struct {
	PaginationFilter
	Service      string `query:"service"`
	ClientID     string `query:"cliendTd"`
	UpdateID     string `query:"updateID"`
	RoomID       string `query:"roomId"`
	Action       string `query:"action"`
	ResponseBody string `query:"responseBody"`
	IsSucceed    *bool  `query:"isSucceed"`
	FromDate     string `query:"fromDate"`
	ToDate       string `query:"toDate"`
}
