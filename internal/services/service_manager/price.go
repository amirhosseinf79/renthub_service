package manager

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (s *sm) asyncPrice(service dto.SiteEntry) (serviceResult dto.ServiceStats) {
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
		Amount: service.Price,
	}

	selectedService, ok := s.apiServices[service.Site]
	if !ok {
		return
	}

	log, err = selectedService.EditPricePerDays(fields)
	s.recordResult(&serviceResult, service.Code, log, err)
	return
}

func (s *sm) PriceUpdate() dto.ManagerResponse {
	var results []dto.ServiceStats
	for _, service := range s.services {
		// go s.asyncPrice(service)
		result := s.asyncPrice(service)
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
