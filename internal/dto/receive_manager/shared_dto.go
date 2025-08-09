package receive_manager_dto

type SiteResponse struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	Response     any    `json:"response,omitempty"`
}

type RecieveResponse map[string]SiteResponse
