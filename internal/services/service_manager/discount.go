package manager

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (s *sm) asyncDiscount(service dto.SiteEntry, discountPercent int) (serviceResult dto.ServiceStats) {
	serviceResult = s.initServiceStatus(service.Site)
	var log *models.Log
	var err error

	fields := dto.UpdateFields{
		RequiredFields: dto.RequiredFields{
			UserID:   s.userID,
			ClientID: s.responseHead.ClientID,
		},
		RoomID: service.Code,
		Dates:  s.dates,
		Amount: discountPercent,
	}

	selectedService, ok := s.apiServices[service.Site]
	if !ok {
		return
	}

	switch discountPercent {
	case 0:
		log, err = selectedService.RemoveDiscount(fields)
	default:
		log, err = selectedService.AddDiscount(fields)
	}
	s.recordResult(&serviceResult, service.Code, log, err)
	return
}

func (s *sm) DiscountUpdate(discountPercent int) dto.ManagerResponse {
	var results []dto.ServiceStats
	for _, service := range s.services {
		// go s.asyncDiscount(service, discountPercent)
		result := s.asyncDiscount(service, discountPercent)
		results = append(results, result)
	}

	result := dto.ManagerResponse{
		ReqHeaderEntry: s.responseHead,
		OveralStatus:   "operating",
		Results:        results,
	}
	s.tryWebHook(result)
	return result
}
