package interfaces

type FetchService interface {
	RequestBody(bodyRow any) error
	RequestQuery(queryRow any) error
	CommitRequest() error
	ParseBody(response, err any) error
	Start(body any, contentType string) error
	ParseInterface(ApiResponseManager) error
	PrintRequestDump()
	Json(any) error
	SetHeaders()
	Ok() bool
}
