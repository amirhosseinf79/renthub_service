package dto

type ManagerConfig struct {
	SendWebHookSeperately bool
}

type ServiceStats struct {
	Site         string `json:"site"`
	Code         string `json:"code"`
	Status       string `json:"status"`
	ErrorMessage string `json:"errorMessage"`
}

type ManagerResponse struct {
	ReqHeaderEntry
	OveralStatus string         `json:"overalStatus"`
	Results      []ServiceStats `json:"results"`
}

func (m *ManagerResponse) SetOveralStatus() {
	counter := 0
	for _, result := range m.Results {
		if result.Status == "success" {
			counter += 1
		}
	}
	if len(m.Results) > 0 {
		switch counter {
		case 0:
			m.OveralStatus = "failed"
		case len(m.Results):
			m.OveralStatus = "success"
		default:
			m.OveralStatus = "partial_success"
		}
	} else {
		m.OveralStatus = "operating"
	}

}
