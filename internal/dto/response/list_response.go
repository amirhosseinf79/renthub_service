package response_dto

import "math"

type ListResponse[T any] struct {
	TotalCount int64 `json:"totalCount"`
	Items      []T   `json:"items"`
	Meta       Meta  `json:"metaData"`
}

func NewListResponse[T any](total int64, page, pagesize uint, list []T) *ListResponse[T] {
	endPage := uint(math.Ceil(float64(total) / float64(pagesize)))
	var nextPage uint = 0
	if page < endPage {
		nextPage = page + 1
	}
	return &ListResponse[T]{
		TotalCount: total,
		Items:      list,
		Meta: Meta{
			CurrentPage: page,
			NextPage:    nextPage,
			EndPage:     endPage,
		},
	}
}
