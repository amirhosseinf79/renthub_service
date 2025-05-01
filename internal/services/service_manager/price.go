package manager

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (s *sm) asyncPrice(service dto.SiteEntry) {
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
		Amount: service.Price,
	}

	selectedService, ok := s.apiServices[service.Site]
	if !ok {
		return
	}

	log, err = selectedService.EditPricePerDays(fields)
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

func (s *sm) PriceUpdate() dto.ManagerResponse {
	for _, service := range s.services {
		go s.asyncPrice(service)
	}

	result := dto.ManagerResponse{
		ReqHeaderEntry: s.responseHead,
		OveralStatus:   "operating",
	}
	return result
}
