package requests

import (
	"encoding/json"
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
	if contentType == "body" {
		err = f.requestBody(body)
	}
	if err != nil {
		return err
	}
	f.setHeaders()
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
		result = dto.ErrInvalidRequest
	}
	if f.httpResp.StatusCode == 401 {
		result = dto.ErrorUnauthorized
	} else if f.httpResp.StatusCode == 403 {
		result = dto.ErrorPermission
	}
	return ok, result
}
