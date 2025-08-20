package jabama_dto

type CalendarRoomResponse struct {
	Result  CalendarRoomResult `json:"result"`
	Success bool               `json:"success"`
}

type CalendarRoomResult struct {
	Calendar       []RoomCalendar      `json:"calendar"`
	Accommodations []RoomAccommodation `json:"accommodations"`
	Filters        any                 `json:"filters"`
	CalendarBanner any                 `json:"calendar_banner"`
	CalendarHint   any                 `json:"calendar_hint"`
	SmartPrice     bool                `json:"smart_price"`
	Guarantee      bool                `json:"guarantee"`
}

type RoomCalendar struct {
	JalaliDateString  string  `json:"jalaliDateString"`
	Date              string  `json:"date"`
	Status            string  `json:"status"`
	Type              string  `json:"type"`
	Price             int64   `json:"price"`
	Discount          int64   `json:"discount"`
	DiscountedPrice   int64   `json:"discountedPrice"`
	IsHoliday         bool    `json:"isHoliday"`
	IsWeekend         bool    `json:"isWeekend"`
	IsCustomHoliday   bool    `json:"isCustomHoliday"`
	MaxAvailableUnits int64   `json:"maxAvailableUnits"`
	AvailableUnits    int64   `json:"availableUnits"`
	MinNight          int64   `json:"minNight"`
	RecommendedPrice  int64   `json:"recommendedPrice"`
	CanPricing        bool    `json:"canPricing"`
	NeedToBeUpdated   bool    `json:"needToBeUpdated"`
	IsPeak            bool    `json:"isPeak"`
	IsPackaged        bool    `json:"isPackaged"`
	PackageColor      string  `json:"packageColor"`
	DemandType        string  `json:"demandType"`
	FilterType        string  `json:"filterType"`
	Hint              *string `json:"hint"`
	SmartPrice        int64   `json:"smartPrice"`
	Guarantee         bool    `json:"guarantee"`
	HasSmartPrice     bool    `json:"hasSmartPrice"`
}

type RoomAccommodation struct {
	ID                   string `json:"id"`
	ComplexID            string `json:"complexId"`
	Code                 int64  `json:"code"`
	UnitCount            int64  `json:"unitCount"`
	Title                string `json:"title"`
	LastUpdateText       string `json:"lastUpdateText"`
	Status               string `json:"status"`
	StatusText           string `json:"statusText"`
	NeedToBeUpdated      bool   `json:"needToBeUpdated"`
	SellableHintText     string `json:"sellableHintText"`
	IsSellable           bool   `json:"isSellable"`
	IsComplex            bool   `json:"isComplex"`
	PlaceType            string `json:"placeType"`
	AffiliateLink        string `json:"affiliateLink"`
	AffiliateDescription string `json:"affiliateDescription"`
	SmartPricing         bool   `json:"smartPricing"`
	ReservationType      string `json:"reservationType"`
}
