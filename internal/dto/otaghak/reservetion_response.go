package otaghak_dto

type ReservationResponse struct {
	Count    int                 `json:"count"`
	Requests []ReservationDetail `json:"requests"`
}

func (r *ReservationResponse) GetList() []ReservationDetail {
	return r.Requests
}

type ReservationDetail struct {
	ID                       int      `json:"id"`
	BookingID                int      `json:"bookingId"`
	BookingCode              string   `json:"bookingCode"`
	BookingStatus            string   `json:"bookingStatus"`
	FromDateTime             string   `json:"fromDateTime"`
	ToDateTime               string   `json:"toDateTime"`
	PersonCount              int      `json:"personCount"`
	ExtraPersonCount         int      `json:"extraPersonCount"`
	HostDeadLineTime         int      `json:"hostDeadLineTime"`
	HostDeadLineTimeTotalSec int      `json:"hostDeadLineTimeTotalSecond"`
	GuestPaymentTime         int      `json:"guestPaymentTime"`
	GuestPaymentTimeTotalSec int      `json:"guestPaymentTimeTotalSecond"`
	RoomID                   int      `json:"roomId"`
	RoomTitle                string   `json:"roomTitle"`
	RoomImageID              int      `json:"roomImageId"`
	RoomAddress              string   `json:"roomAddress"`
	CityFaName               string   `json:"cityFaName"`
	RoomTypeName             string   `json:"roomTypeName"`
	RoomRate                 float64  `json:"roomRate"`
	GuestID                  int      `json:"guestId"`
	GuestFirstName           string   `json:"guestFirstName"`
	GuestLastName            string   `json:"guestLastName"`
	GuestMobileNumber        string   `json:"guestMobileNumber"`
	GuestEmail               *string  `json:"guestEmail"`
	HostID                   int      `json:"hostId"`
	HostFirstName            string   `json:"hostFirstName"`
	HostLastName             string   `json:"hostLastName"`
	StayingTime              int      `json:"stayingTime"`
	TotalAmount              float64  `json:"totalAmount"`
	HostShareAmount          float64  `json:"hostShareAmount"`
	DiscountAmount           float64  `json:"discountAmount"`
	GrossAmount              float64  `json:"grossAmount"`
	ExtraPersonAmount        float64  `json:"extraPersonAmount"`
	IsInstantBooking         bool     `json:"isInstantBooking"`
	TotalPaymentAmount       float64  `json:"totalPaymentAmount"`
	IsCancelable             bool     `json:"isCancelable"`
	IsAllowSendMessage       bool     `json:"isAllowSendMessage"`
	AllowChatStatus          string   `json:"allowChatStatus"`
	AllowChatStatusDesc      string   `json:"allowChatStatusDescription"`
	IsAllowComment           bool     `json:"isAllowComment"`
	UserPoint                *float64 `json:"userPoint"`
	GuestProfileImageID      int      `json:"guestProfileImageId"`
	AllowReplyComment        bool     `json:"allowReplyComment"`
}
