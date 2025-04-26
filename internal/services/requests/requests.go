package requests

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/google/go-querystring/query"
)

type fetchS struct {
	url      string
	method   string
	headers  map[string]string
	extra    map[string]string
	httpReq  *http.Request
	httpResp *http.Response
}

func New(method, url string, headers, extra map[string]string) interfaces.FetchService {
	return &fetchS{
		method:  method,
		url:     url,
		headers: headers,
		extra:   extra,
	}
}

func (f *fetchS) RequestBody(bodyRow any) error {
	body, err := json.Marshal(bodyRow)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(f.method, f.url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	for k, v := range f.headers {
		value, ok := f.extra[k]
		if ok {
			v = fmt.Sprintf(v, value)
		}
		req.Header.Set(k, v)
	}
	f.httpReq = req
	return nil
}

func (f *fetchS) RequestQuery(queryRow any) error {
	v, err := query.Values(queryRow)
	if err != nil {
		return err
	}
	fullURL := fmt.Sprintf("%s?%s", f.url, v.Encode())

	req, err := http.NewRequest(f.method, fullURL, nil)
	if err != nil {
		return err
	}

	for k, v := range f.headers {
		value, ok := f.extra[k]
		if ok {
			v = fmt.Sprintf(v, value)
		}
		req.Header.Set(k, v)
	}
	f.httpReq = req
	return nil
}

func (f *fetchS) PrintRequestDump() {
	dump, err := httputil.DumpRequestOut(f.httpReq, true)
	if err != nil {
		fmt.Println("Error dumping request:", err)
	}
	fmt.Println(string(dump))
}

func (f *fetchS) CommitRequest() (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(f.httpReq)
	if err != nil {
		return nil, err
	}
	f.httpResp = resp
	return resp, nil
}

func (f *fetchS) Unzip() ([]byte, error) {
	bodyReader, err := gzip.NewReader(f.httpResp.Body)
	if err != nil {
		return nil, err
	}
	defer bodyReader.Close()
	body, err := io.ReadAll(bodyReader)
	return body, err
}

func (f *fetchS) Ok() bool {
	return 200 <= f.httpResp.StatusCode && f.httpResp.StatusCode < 300
}
