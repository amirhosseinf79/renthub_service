package manager

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (s *sm) asyncDiscount(service dto.SiteEntry, discountPercent int) {
	var results []dto.ServiceStats
	serviceResult := s.initServiceStatus(service.Site)
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

	serviceResult.Code = service.Code
	s.logger.RecordLog(log)
	if err != nil {
		serviceResult.Status = "failed"
		serviceResult.ErrorMessage = err.Error()
	}
	results = append(results, serviceResult)
	result := dto.ManagerResponse{
		ReqHeaderEntry: s.responseHead,
		Results:        results,
	}
	result.SetOveralStatus()
	s.tryWebHook(result)
}

func (s *sm) DiscountUpdate(discountPercent int) dto.ManagerResponse {
	for _, service := range s.services {
		go s.asyncDiscount(service, discountPercent)
	}

	result := dto.ManagerResponse{
		ReqHeaderEntry: s.responseHead,
		OveralStatus:   "operating",
	}
	return result
}
