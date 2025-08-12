package shab_dto

type ReservationResponse struct {
	Data ReservationData `json:"data"`
	Meta Meta            `json:"meta"`
}

func (r *ReservationResponse) GetList() []ReservationRecord {
	return r.Data.Records
}

type ReservationData struct {
	MissedReserves int                 `json:"missed_reserves"`
	Records        []ReservationRecord `json:"records"`
	Pagination     Pagination          `json:"pagination"`
}

type ReservationRecord struct {
	ReserveID           string        `json:"reserve_id"`
	IsThirdParty        bool          `json:"is_thirdparty"`
	CreatedAt           string        `json:"created_at"`
	CreatedAtTimestamp  int64         `json:"created_at_timestamp"`
	GuestsCount         int           `json:"guests_count"`
	UnseenMessagesCount int           `json:"unseen_messages_count"`
	CheckinToCheckout   string        `json:"checkin_to_checkout"`
	CheckinTimestamp    int64         `json:"checkin_timestamp"`
	UserGuest           User          `json:"user_guest"`
	Calendar            []interface{} `json:"calendar"`
	UserHost            User          `json:"user_host"`
	Checkin             string        `json:"checkin"`
	Checkout            string        `json:"checkout"`
	Duration            int           `json:"duration"`
	ExpiredAt           string        `json:"expired_at"`
	CheckoutTimestamp   int64         `json:"checkout_timestamp"`
	MaxCheckinTime      *string       `json:"max_checkin_time"`
	CheckinTime         string        `json:"checkin_time"`
	CheckoutTime        string        `json:"checkout_time"`
	ExpiredAtTimestamp  int64         `json:"expired_at_timestamp"`
	ExtraPersonCount    int           `json:"extra_person_count"`
	NormalPersonCount   int           `json:"normal_person_count"`
	Status              Status        `json:"status"`
	LastState           string        `json:"last_state"`
	StatusText          string        `json:"status_text"`
	UnsuccessfulText    string        `json:"unsuccessful_text"`
	HasHostReview       bool          `json:"has_host_review"`
	IsInstant           bool          `json:"is_instant"`
	IsMidTerm           int           `json:"is_mid_term"`
	Invoice             Invoice       `json:"invoice"`
	House               House         `json:"house"`
}

type User struct {
	Fullname string  `json:"fullname"`
	Picture  string  `json:"picture"`
	Mobile   *string `json:"mobile"`
	ID       string  `json:"id"`
}

type Status struct {
	State string `json:"state"`
	Actor string `json:"actor"`
}

type Invoice struct {
	Price Price  `json:"price"`
	ID    string `json:"id"`
}

type Price struct {
	Amount            int    `json:"amount"`
	CurrencyName      string `json:"currency_name"`
	CurrencyShortForm string `json:"currency_short_form"`
}

type House struct {
	CoverPhoto string   `json:"cover_photo"`
	RoomsCount int      `json:"rooms_count"`
	Location   Location `json:"location"`
	Province   string   `json:"province"`
	Title      string   `json:"title"`
	Rules      Rules    `json:"rules"`
	City       string   `json:"city"`
	ID         int      `json:"id"`
}

type Location struct {
	Address  *string         `json:"address"`
	Location *GeoCoordinates `json:"location"`
}

type GeoCoordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Rules struct {
	Pets             *string `json:"pets"`
	Ceremony         *string `json:"ceremony"`
	CanChangeDate    int     `json:"can_change_date"`
	Documents        string  `json:"documents"`
	CancellationPlan int     `json:"cancellation_plan"`
}

type Pagination struct {
	CurrentPage int `json:"current_page"`
	Total       int `json:"total"`
}

type Meta struct {
	Status   int      `json:"status"`
	Messages []string `json:"messages"`
}
