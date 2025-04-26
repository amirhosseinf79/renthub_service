package dto

type RequiredFields struct {
	UserID   uint //`json:"user_id" form:"user_id" query:"user_id"`
	ClientID string
}

type UpdateFields struct {
	RequiredFields
	RoomID string
	Amount uint
}

type ApiEasyLogin struct {
	RequiredFields
	Username string
	Password string
}

type ApiEndpoints struct {
	LoginFirstStep  string
	LoginSecondStep string
	LoginWithPass   string
	GetProfile      string
	OpenCalendar    string
	CloseCalendar   string
	EditPricePerDay string
	AddDiscount     string
	RemoveDiscount  string
	SetMinNight     string
	UnsetMinNight   string
}
