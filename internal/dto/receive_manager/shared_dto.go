package receive_manager_dto

type SiteResponse struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	ResponseCode int    `json:"responseCode"`
	Response     any    `json:"response,omitempty"`
}

type RecieveResponse map[string]SiteResponse
