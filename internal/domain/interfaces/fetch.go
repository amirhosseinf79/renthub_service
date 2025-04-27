package interfaces

type FetchService interface {
	Start(body any, contentType string) error
	ParseInterface(ApiResponseManager) error
	Ok() (bool, error)
}
