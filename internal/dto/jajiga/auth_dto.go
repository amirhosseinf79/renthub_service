package jajiga_dto

type JajigaAuthRequestBody struct {
	Mobile     string `json:"mobile"`
	Password   string `json:"password"`
	ISO2       string `json:"iso2"`
	ClientID   string `json:"client_id"`
	ClientType string `json:"client_type"`
}
