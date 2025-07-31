package response_dto

type Meta struct {
	CurrentPage uint `json:"currentPage"`
	NextPage    uint `json:"nextPage"`
	EndPage     uint `json:"endPage"`
}
