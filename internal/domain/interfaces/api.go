package interfaces

import (
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

type ApiService interface {
	EasyLogin(dto.ApiEasyLogin) error
	SendOtp(dto.RequiredFields, string) error
	VerifyOtp(dto.RequiredFields, string) error
	CheckLogin(dto.RequiredFields) error
	OpenCalendar(dto.UpdateFields) error
	CloseCalendar(dto.UpdateFields) error
	EditPricePerDays(dto.UpdateFields) error
	AddDiscount(dto.UpdateFields) error
	RemoveDiscount(dto.UpdateFields) error
	SetMinNight(dto.UpdateFields) error
	UnsetMiniNight(dto.UpdateFields) error
	// GetRooms()
	// GetRoomDetails()
}
