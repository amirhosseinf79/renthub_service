package interfaces

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

type ApiService interface {
	Set(service string) ApiService
	EasyLogin(dto.ApiEasyLogin) *models.Log
	SendOtp(dto.RequiredFields, string) *models.Log
	VerifyOtp(dto.RequiredFields, string) *models.Log
	CheckLogin(dto.RequiredFields) *models.Log
	OpenCalendar(dto.UpdateFields) *models.Log
	CloseCalendar(dto.UpdateFields) *models.Log
	EditPricePerDays(dto.UpdateFields) *models.Log
	AddDiscount(dto.UpdateFields) *models.Log
	RemoveDiscount(dto.UpdateFields) *models.Log
	SetMinNight(dto.UpdateFields) *models.Log
	UnsetMiniNight(dto.UpdateFields) *models.Log
	RecordLog(*models.Log) error
	// GetRooms()
	// GetRoomDetails()
}

type ApiResponseManager interface {
	GetResult() (bool, string)
	GetToken() *models.ApiAuth
}
