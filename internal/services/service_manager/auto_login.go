package manager

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (s *sm) asyncAutoLogin(service dto.SiteEntry, chResult chan dto.ServiceStats) {
	serviceResult := s.initServiceStatus(service.Site)
	var log *models.Log
	var err error

	fields := dto.RequiredFields{
		UserID:   s.userID,
		ClientID: s.responseHead.ClientID,
	}
	selectedService, ok := s.apiServices[service.Site]
	if !ok {
		return
	}

	log, err = selectedService.AutoLogin(fields)
	s.recordResult(&serviceResult, service.Code, log, err)
	chResult <- serviceResult
}

func (s *sm) ManageAutoLogin() dto.ManagerResponse {
	chResult := make(chan dto.ServiceStats)
	var results []dto.ServiceStats

	for _, service := range s.services {
		go s.asyncAutoLogin(service, chResult)
	}

	for range len(s.services) {
		results = append(results, <-chResult)
	}
	close(chResult)

	result := dto.ManagerResponse{
		ReqHeaderEntry: s.responseHead,
		Results:        results,
	}
	result.SetOveralStatus()
	return result
}
