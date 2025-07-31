package dto

type PaginationFilter struct {
	Page     uint `query:"page"`
	PageSize uint `query:"pageSize"`
}
