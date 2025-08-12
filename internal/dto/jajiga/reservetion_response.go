package jajiga_dto

type ReservationResponse struct {
	Pagination Pagination `json:"pagination"`
	Sections   []Section  `json:"sections"`
	Items      []any      `json:"items"`
}

func (r *ReservationResponse) GetList() any {
	return r.Items
}

type Pagination struct {
	Total   int `json:"total"`
	Count   int `json:"count"`
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

type Section struct {
	Name   string `json:"name"`
	Count  int    `json:"count"`
	Active bool   `json:"active"`
}
