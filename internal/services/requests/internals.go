package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

func (f *fetchS) requestBody(bodyRow any) error {
	body, err := json.Marshal(bodyRow)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(f.method, f.url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	if f.method == "GET" {
		req.ContentLength = -1
	}
	f.httpReq = req
	return nil
}

// func (f *fetchS) requestBodyString() error {
// 	payload := strings.NewReader("{\"mobile\": \"09334429096\",\"code\": \"\"}")
// 	req, err := http.NewRequest(f.method, f.url, payload)
// 	if err != nil {
// 		return err
// 	}
// 	f.httpReq = req
// 	return nil
// }

// func (f *fetchS) requestQuery(queryRow any) error {
// 	v, err := query.Values(queryRow)
// 	if err != nil {
// 		return err
// 	}
// 	fullURL := fmt.Sprintf("%s?%s", f.url, v.Encode())

// 	req, err := http.NewRequest(f.method, fullURL, nil)
// 	if err != nil {
// 		return err
// 	}
// 	f.httpReq = req
// 	return nil
// }

func (f *fetchS) setHeaders() {
	for k, v := range f.headers {
		hasVar := bytes.Contains([]byte(v), []byte("%v"))
		value, ok := f.extra[k]
		if ok && hasVar {
			v = fmt.Sprintf(v, value)
		} else if hasVar {
			v = fmt.Sprintf(v, "")
		}
		f.httpReq.Header.Set(k, v)
	}
}

func (f *fetchS) dumpRequest() {
	dump, err := httputil.DumpRequestOut(f.httpReq, true)
	if err != nil {
		fmt.Println("Error dumping request:", err)
	}
	dumps := string(dump)
	f.logger.RequestBody = dumps
	fmt.Println(dumps)
}

func (f *fetchS) commitRequest() error {
	client := &http.Client{}
	resp, err := client.Do(f.httpReq)
	if err != nil {
		return err
	}
	f.logger.StatusCode = resp.StatusCode
	fmt.Println("Status:", resp.StatusCode)
	f.httpResp = resp
	return nil
}

func (f *fetchS) parseBodyResponse() error {
	bodyBytes, err := f.readBodyResponse()
	if err != nil {
		return err
	}
	f.httpResp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	f.logger.ResponseBody = string(bodyBytes)
	return nil
}

func (f *fetchS) readBodyResponse() ([]byte, error) {
	body, err := io.ReadAll(f.httpResp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
