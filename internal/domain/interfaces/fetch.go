package interfaces

import "net/http"

type FetchService interface {
	RequestBody(bodyRow any) error
	RequestQuery(queryRow any) error
	CommitRequest() (*http.Response, error)
	PrintRequestDump()
	Unzip() ([]byte, error)
	Ok() bool
}
