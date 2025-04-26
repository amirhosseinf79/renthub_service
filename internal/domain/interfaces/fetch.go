package interfaces

type FetchService interface {
	RequestBody(bodyRow any) error
	RequestQuery(queryRow any) error
	CommitRequest() error
	PrintRequestDump()
	Json(any) error
	Ok() bool
}
