package jabama_dto

type ReservationResponse struct {
	Result  ReservationResult `json:"result"`
	Success bool              `json:"success"`
}

func (r *ReservationResponse) GetList() []ReservationData {
	return r.Result.Data
}

type ReservationResult struct {
	Count                     int                       `json:"count"`
	Data                      []ReservationData         `json:"data"`
	StatusList                []ReservationStatus       `json:"status_list"`
	Filter                    ReservationFilter         `json:"filter"`
	SortList                  []ReservationSort         `json:"sort_list"`
	NotSuppliedGovernanceHint NotSuppliedGovernanceHint `json:"not_supplied_governance_hint"`
}

type ReservationData struct {
	Tag               string                  `json:"tag"`
	Color             string                  `json:"color"`
	AccommodationName string                  `json:"accommodation_name"`
	Order             ReservationOrder        `json:"order"`
	Guest             ReservationGuest        `json:"guest"`
	Capacity          ReservationCapacity     `json:"capacity"`
	Check             ReservationCheck        `json:"check"`
	Price             int64                   `json:"price"`
	NewPrice          int64                   `json:"new_price"`
	UnitTitle         *string                 `json:"unit_title"`
	ReservationCode   string                  `json:"reservation_code"`
	RemainingTime     *string                 `json:"remaining_time"`
	GuestReview       *ReservationGuestReview `json:"guest_review"`
	Actions           []ReservationAction     `json:"actions"`
	Options           []ReservationOption     `json:"options"`
	NsReasonTitle     *string                 `json:"ns_reason_title"`
}

type ReservationOrder struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type ReservationGuest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type ReservationCapacity struct {
	Base  int `json:"base"`
	Extra int `json:"extra"`
}

type ReservationCheck struct {
	In         ReservationCheckDate `json:"in"`
	Out        ReservationCheckDate `json:"out"`
	Text       string               `json:"text"`
	InChange   ReservationCheckDate `json:"in_change"`
	OutChange  ReservationCheckDate `json:"out_change"`
	ChangeText string               `json:"change_text"`
}

type ReservationCheckDate struct {
	Day     int    `json:"day"`
	Month   string `json:"month"`
	WeekDay string `json:"week_day"`
}

type ReservationGuestReview struct {
	ReviewID *string `json:"review_id"`
	Score    *int    `json:"score"`
	Status   *string `json:"status"`
}

type ReservationAction struct {
	Type   string `json:"type"`
	Active bool   `json:"active"`
}

type ReservationOption struct {
	Type   string `json:"type"`
	Active bool   `json:"active"`
}

type ReservationStatus struct {
	Name       string `json:"name"`
	NameFa     string `json:"name_fa"`
	IsSelected bool   `json:"is_selected"`
}

type ReservationFilter struct {
	Accommodations []ReservationAccommodation `json:"accommodations"`
}

type ReservationAccommodation struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ReservationSort struct {
	Key      string `json:"key"`
	NameFa   string `json:"name_fa"`
	Selected bool   `json:"selected"`
}

type NotSuppliedGovernanceHint struct {
	Title       string                        `json:"title"`
	Description string                        `json:"description"`
	Color       string                        `json:"color"`
	Icon        string                        `json:"icon"`
	Actions     []NotSuppliedGovernanceAction `json:"actions"`
	Details     []NotSuppliedGovernanceDetail `json:"details"`
}

type NotSuppliedGovernanceAction struct {
	Type  string `json:"type"`
	Title string `json:"title"`
}

type NotSuppliedGovernanceDetail struct {
	AccommodationTitle string `json:"accommodation_title"`
	Fine               string `json:"fine"`
	NextStep           string `json:"next_step"`
}
