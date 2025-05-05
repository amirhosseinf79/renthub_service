package interfaces

type FetchService interface {
	Start(body any, contentType string) error
	ParseInterface(any) error
	Ok() (bool, error)
}
