package update_manager_v1

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	request_v1 "github.com/amirhosseinf79/renthub_service/internal/dto/request/v1"
)

func (s *sm) asyncPrice(service request_v1.SiteEntry, chResult chan request_v1.ServiceStats) {
	serviceResult := s.initServiceStatus(service.Site)
	var log *models.Log
	var err error

	fields := dto.UpdateFields{
		RequiredFields: dto.RequiredFields{
			UserID:   s.userID,
			ClientID: s.requestHeader.ClientID,
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
	chResult <- serviceResult
}

func (s *sm) PriceUpdate() request_v1.ManagerResponse {
	chResult := make(chan request_v1.ServiceStats)
	var results []request_v1.ServiceStats

	for _, service := range s.services {
		go s.asyncPrice(service, chResult)
	}

	for range len(s.services) {
		results = append(results, <-chResult)
	}
	close(chResult)

	result := request_v1.ManagerResponse{
		ReqHeaderEntry: s.requestHeader,
		Results:        results,
	}
	result.SetOveralStatus()
	return result
}
