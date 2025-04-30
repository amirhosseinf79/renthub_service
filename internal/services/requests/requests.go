package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

type fetchS struct {
	url      string
	method   string
	headers  map[string]string
	extra    map[string]string
	httpReq  *http.Request
	httpResp *http.Response
	logger   *models.Log
}

func New(method, url string, headers, extra map[string]string, logger *models.Log) interfaces.FetchService {
	fmt.Println("Method:", method)
	fmt.Println("Full URL:", url)
	logger.RequestURL = url

	return &fetchS{
		method:  method,
		url:     url,
		headers: headers,
		extra:   extra,
		logger:  logger,
	}
}

func (f *fetchS) Start(body any, contentType string) error {
	var err error
	payload := &bytes.Buffer{}
	headerCType := ""
	switch contentType {
	case "body":
		payload, err = f.requestBody(body)
		headerCType = "application/json; charset=UTF-8"
	case "query":
		err = f.requestQuery(body)
	case "multipart":
		payload, headerCType, err = f.requestMultipart(body)
	default:
		err = fmt.Errorf("unsupported content type: %s", contentType)
	}
	if err != nil {
		return err
	}
	f.logger.RequestBody = payload.String()
	err = f.NewRequest(payload)
	if err != nil {
		return err
	}
	f.setHeaders(headerCType)
	if f.logger != nil {
		f.dumpRequest()
	}
	err = f.commitRequest()
	if err != nil {
		return err
	}
	return nil
}

func (f *fetchS) ParseInterface(response interfaces.ApiResponseManager) (err error) {
	if f.logger != nil {
		err := f.parseBodyResponse()
		if err != nil {
			return err
		}
	}
	body, err := f.readBodyResponse()
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, response)
	if err != nil {
		return
	}
	return
}

func (f *fetchS) Ok() (bool, error) {
	ok := 200 <= f.httpResp.StatusCode && f.httpResp.StatusCode < 300
	var result error
	if !ok {
		if f.httpResp.StatusCode == 401 {
			result = dto.ErrorUnauthorized
		} else if f.httpResp.StatusCode == 403 {
			result = dto.ErrorPermission
		} else if f.httpResp.StatusCode == 404 {
			result = dto.ErrRoomNotFound
		} else {
			result = dto.ErrInvalidRequest
		}
	}
	return ok, result
}
