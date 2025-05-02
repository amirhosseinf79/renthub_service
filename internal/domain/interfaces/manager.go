package interfaces

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

type ServiceManager interface {
	SetConfigs(userID uint, header dto.ReqHeaderEntry, services []dto.SiteEntry, dates []string) ServiceManager
	CalendarUpdate(action string) dto.ManagerResponse
	MinNightUpdate(limitDays int) dto.ManagerResponse
	DiscountUpdate(discountPercent int) dto.ManagerResponse
	SendWebhook(response dto.ManagerResponse) (*models.Log, error)
	PriceUpdate() dto.ManagerResponse
	CheckAuth() dto.ManagerResponse
}
