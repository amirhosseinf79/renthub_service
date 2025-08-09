package dto

type RequiredFields struct {
	UserID   uint //`json:"user_id" form:"user_id" query:"user_id"`
	ClientID string
}

type GetDetail struct {
	RequiredFields
	RoomID string
	Page   int
}

type UpdateFields struct {
	RequiredFields
	RoomID string
	Dates  []string
	Amount int
}

type SiteFilters = map[string]any

type RecieveFields struct {
	RequiredFields
	Filters SiteFilters
}

type ApiEasyLogin struct {
	RequiredFields
	Username string
	Password string
}

type EndP struct {
	Address     string
	Method      string
	ContentType string
}

type ApiEndpoints struct {
	LoginFirstStep     EndP
	LoginSecondStep    EndP
	LoginWithPass      EndP
	GetProfile         EndP
	OpenCalendar       EndP
	CloseCalendar      EndP
	EditPricePerDay    EndP
	AddDiscount        EndP
	RemoveDiscount     EndP
	SetMinNight        EndP
	UnsetMinNight      EndP
	GETRooms           EndP
	GETRoomDetails     EndP
	GetCalendarDetails EndP
	GetReservations    EndP
}

type ApiSettings struct {
	ApiURL    string
	Endpoints ApiEndpoints
	Headers   map[string]string
}

type ApiErrResponse struct {
	Code    string
	Message string
}
