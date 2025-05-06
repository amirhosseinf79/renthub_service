package interfaces

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

type ApiService interface {
	AutoLogin(dto.RequiredFields) (*models.Log, error)
	SendOtp(dto.RequiredFields, string) (*models.Log, error)
	VerifyOtp(dto.RequiredFields, dto.OTPCreds) (*models.Log, error)
	CheckLogin(dto.RequiredFields) (*models.Log, error)
	OpenCalendar(dto.UpdateFields) (*models.Log, error)
	CloseCalendar(dto.UpdateFields) (*models.Log, error)
	EditPricePerDays(dto.UpdateFields) (*models.Log, error)
	AddDiscount(dto.UpdateFields) (*models.Log, error)
	RemoveDiscount(dto.UpdateFields) (*models.Log, error)
	SetMinNight(dto.UpdateFields) (*models.Log, error)
	UnsetMiniNight(dto.UpdateFields) (*models.Log, error)
	// GetRooms()
	// GetRoomDetails()
}

type ApiResponseManager interface {
	GetResult() (bool, string)
	GetToken() *models.ApiAuth
}
