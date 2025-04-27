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

type EndP struct {
	Address     string
	Method      string
	ContentType string
}

type ApiEndpoints struct {
	LoginFirstStep  EndP
	LoginSecondStep EndP
	LoginWithPass   EndP
	GetProfile      EndP
	OpenCalendar    EndP
	CloseCalendar   EndP
	EditPricePerDay EndP
	AddDiscount     EndP
	RemoveDiscount  EndP
	SetMinNight     EndP
	UnsetMinNight   EndP
}

type ApiSettings struct {
	ApiURL    string
	Endpoints ApiEndpoints
	Headers   map[string]string
}
