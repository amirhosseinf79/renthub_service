package homsa_dto

type ReservationResponse struct {
	Data []ReservationData `json:"data"`
	Meta Meta              `json:"meta"`
}

func (r *ReservationResponse) GetList() any {
	return r.Data
}

type ReservationData struct {
	ID                int               `json:"id"`
	RoomID            int               `json:"room_id"`
	HostID            int               `json:"host_id"`
	GuestID           int               `json:"guest_id"`
	Checkin           string            `json:"checkin"`
	Checkout          string            `json:"checkout"`
	GuestFullName     string            `json:"guest_full_name"`
	GuestProfileImage string            `json:"guest_profile_image"`
	RoomName          string            `json:"room_name"`
	RoomImage         string            `json:"room_image"`
	RoomLocation      string            `json:"room_location"`
	GuestNumber       int               `json:"guest_number"`
	GuestNumberUnit   string            `json:"guest_number_unit"`
	Nights            string            `json:"nights"`
	ReserveDuration   string            `json:"reserve_duration"`
	SecondsToExpire   int               `json:"seconds_to_expire"`
	Discount          string            `json:"discount"`
	HostAmount        string            `json:"host_amount"`
	TotalAmount       string            `json:"total_amount"`
	Status            ReservationStatus `json:"status"`
	GuestPhoneNumber  string            `json:"guest_phone_number,omitempty"`
}

type ReservationStatus struct {
	Slug       string `json:"slug"`
	Title      string `json:"title"`
	TitleColor string `json:"title_color"`
	BgColor    string `json:"bg_color"`
}

type Meta struct {
	CurrentPage int `json:"current_page"`
	From        int `json:"from"`
	LastPage    int `json:"last_page"`
	PerPage     int `json:"per_page"`
	To          int `json:"to"`
	Total       int `json:"total"`
}
