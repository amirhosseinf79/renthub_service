package jabama_dto

type Response struct {
	Result              *authResult `json:"result"`
	TargetURL           *string     `json:"targetUrl"`
	Success             bool        `json:"success"`
	Error               *authError  `json:"error"`
	UnauthorizedRequest bool        `json:"unauthorizedRequest"`
	Wrapped             bool        `json:"__wrapped"`
	TraceID             string      `json:"__traceId"`
}

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
