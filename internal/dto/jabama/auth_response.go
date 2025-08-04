package jabama_dto

import "github.com/amirhosseinf79/renthub_service/internal/domain/models"

type authResult struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	Mobile       string `json:"mobile"`
	NextStep     string `json:"nextStep"`
}

type authError struct {
	ErrorCode        int               `json:"errorCode"`
	Message          string            `json:"message"`
	ServiceMessage   string            `json:"serviceMessage"`
	Details          string            `json:"details"`
	Source           string            `json:"source"`
	ValidationErrors map[string]string `json:"validationErrors"`
}

type Response struct {
	Result              *authResult `json:"result"`
	TargetURL           *string     `json:"targetUrl"`
	Success             bool        `json:"success"`
	Error               *authError  `json:"error"`
	UnauthorizedRequest bool        `json:"unauthorizedRequest"`
	Wrapped             bool        `json:"__wrapped"`
	TraceID             string      `json:"__traceId"`
}

func (h *Response) GetResult() (ok bool, result string) {
	ok = h.Error == nil
	result = "success"
	if !ok {
		result = h.Error.Message
	}
	return ok, result
}

func (h *Response) GetToken() *models.ApiAuth {
	return &models.ApiAuth{
		AccessToken:  h.Result.AccessToken,
		RefreshToken: h.Result.RefreshToken,
	}
}
