package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httputil"

	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/google/go-querystring/query"
)

func (f *fetchS) requestBody(bodyRow any) (*bytes.Buffer, error) {
	body, err := json.Marshal(bodyRow)
	if err != nil {
		return nil, err
	}
	payload := bytes.NewBuffer(body)
	return payload, nil
}

func (f *fetchS) requestMultipart(bodyRow any) (*bytes.Buffer, string, error) {
	bodyS, ok := bodyRow.([]byte)
	if !ok {
		return nil, "", dto.ErrInvalidRequest
	}
	var mapBody map[string]string
	err := json.Unmarshal([]byte(bodyS), &mapBody)
	if err != nil {
		return nil, "", dto.ErrInvalidRequest
	}
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	for key, val := range mapBody {
		err := writer.WriteField(key, val)
		if err != nil {
			return nil, "", err
		}
	}
	cType := writer.FormDataContentType()
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}

	return payload, cType, err
}

// func (f *fetchS) requestUnlenCoded(body any) (*bytes.Buffer, error) {
// 	mapBody, ok := body.(map[string]string)
// 	if !ok {
// 		return nil, dto.ErrInvalidRequest
// 	}
// 	var dates []string
// 	for key, val := range mapBody {
// 		dates = append(dates, fmt.Sprintf("%v=%v", key, val))
// 	}
// 	encoded := bytes.NewBufferString(strings.Join(dates, "&"))
// 	return encoded, nil
// }

func (f *fetchS) requestQuery(queryRow any) error {
	v, err := query.Values(queryRow)
	if err != nil {
		return err
	}
	fullURL := fmt.Sprintf("%s?%s", f.url, v.Encode())
	f.url = fullURL
	return nil
}

func (f *fetchS) NewRequest(body *bytes.Buffer) error {
	req, err := http.NewRequest(f.method, f.url, body)
	if err != nil {
		return err
	}
	if f.method == "GET" {
		req.ContentLength = -1
	}
	f.httpReq = req
	return nil
}

func (f *fetchS) setHeaders(contentType string) {
	if contentType != "" {
		f.httpReq.Header.Set("Content-Type", contentType)
	}
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
	// fmt.Println(dumps)
}

func (f *fetchS) commitRequest() error {
	client := &http.Client{}
	resp, err := client.Do(f.httpReq)
	if err != nil {
		return err
	}
	f.logger.StatusCode = resp.StatusCode
	f.httpResp = resp
	if f.logger != nil {
		err := f.parseBodyResponse()
		if err != nil {
			return err
		}
	}
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
