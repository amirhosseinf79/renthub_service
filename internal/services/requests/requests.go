package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
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
	f.httpReq = req
	return nil
}

func (f *fetchS) SetHeaders() {
	for k, v := range f.headers {
		value, ok := f.extra[k]
		if ok {
			v = fmt.Sprintf(v, value)
		}
		f.httpReq.Header.Set(k, v)
	}
}

func (f *fetchS) PrintRequestDump() {
	dump, err := httputil.DumpRequestOut(f.httpReq, true)
	if err != nil {
		fmt.Println("Error dumping request:", err)
	}
	fmt.Println(string(dump))
}

func (f *fetchS) CommitRequest() error {
	client := &http.Client{}
	resp, err := client.Do(f.httpReq)
	if err != nil {
		return err
	}
	f.httpResp = resp
	return nil
}

func (f *fetchS) Json(v any) error {
	body, err := io.ReadAll(f.httpResp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("Raw Response Body: %s\n", string(body))
	err = json.Unmarshal(body, v)
	return err
}

func (f *fetchS) Start(body any, contentType string) error {
	var err error
	if contentType == "body" {
		err = f.RequestBody(body)
	}
	if err != nil {
		return err
	}
	f.SetHeaders()
	f.PrintRequestDump()
	err = f.CommitRequest()
	if err != nil {
		return err
	}
	return nil
}

func (f *fetchS) ParseBody(response, failed any) error {
	ok, _ := f.Ok()
	if !ok {
		err := f.Json(failed)
		if err != nil {
			return err
		}
		return nil
	}

	err := f.Json(response)
	if err != nil {
		return err
	}
	return nil
}

func (f *fetchS) ParseInterface(response interfaces.ApiResponseManager) (err error) {
	err = f.Json(response)
	if err != nil {
		return
	}
	return
}

func (f *fetchS) Ok() (bool, error) {
	ok := 200 <= f.httpResp.StatusCode && f.httpResp.StatusCode < 300
	var result error
	if !ok {
		result = dto.ErrInvalidRequest
	}

	if f.httpResp.StatusCode == 401 {
		result = dto.ErrorUnauthorized
	} else if f.httpResp.StatusCode == 403 {
		result = dto.ErrorPermission
	}
	return ok, result
}
