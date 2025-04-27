package interfaces

type FetchService interface {
	RequestBody(bodyRow any) error
	RequestQuery(queryRow any) error
	CommitRequest() error
	BodyStart(any) error
	ParseBody(response, err any) error
	ParseInterface(ApiResponseManager) error
	PrintRequestDump()
	Json(any) error
	SetHeaders()
	Ok() bool
}
