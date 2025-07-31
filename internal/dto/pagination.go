package dto

type PaginationFilter struct {
	Page     uint `query:"page" validate:"gt=0"`
	PageSize uint `query:"pageSize" validate:"gt=0,lte=100"`
}
