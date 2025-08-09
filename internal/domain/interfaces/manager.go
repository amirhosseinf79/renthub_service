package interfaces

import (
	receive_manager_dto "github.com/amirhosseinf79/renthub_service/internal/dto/receive_manager"
	request_v1 "github.com/amirhosseinf79/renthub_service/internal/dto/request/v1"
	request_v2 "github.com/amirhosseinf79/renthub_service/internal/dto/request/v2"
)

type ServiceManager interface {
	ManageAutoLogin() request_v1.ManagerResponse
	SetConfigs(userID uint, header request_v1.ReqHeaderEntry, services []request_v1.SiteEntry, dates []string) ServiceManager
	CalendarUpdate(action string) request_v1.ManagerResponse
	MinNightUpdate(limitDays int) request_v1.ManagerResponse
	DiscountUpdate(discountPercent int) request_v1.ManagerResponse
	PriceUpdate() request_v1.ManagerResponse
	CheckAuth() request_v1.ManagerResponse
}

type ServiceUpdateManagerV2 interface {
	ManageAutoLogin() request_v2.ManagerResponse
	SetConfigs(userID uint, header request_v2.ReqHeaderEntry, services []request_v2.SiteEntry, dates []string) ServiceUpdateManagerV2
	CalendarUpdate(action string) request_v2.ManagerResponse
	MinNightUpdate(limitDays int) request_v2.ManagerResponse
	DiscountUpdate(discountPercent int) request_v2.ManagerResponse
	PriceUpdate() request_v2.ManagerResponse
	CheckAuth() request_v2.ManagerResponse
}

type ServiceRecieveManagerV2 interface {
	SetConfigs(userID uint, header request_v2.ReqHeaderEntry, services []request_v2.SiteRecieve) ServiceRecieveManagerV2
	GetReservations() receive_manager_dto.RecieveResponse
}
