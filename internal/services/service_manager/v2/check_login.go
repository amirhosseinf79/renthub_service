package manager_v2

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	request_v2 "github.com/amirhosseinf79/renthub_service/internal/dto/request/v2"
)

func (s *sm) asyncCheckAuth(service request_v2.SiteEntry, chResult chan request_v2.ServiceStats) {
	serviceResult := s.initServiceStatus(service.Site)
	var log *models.Log
	var err error

	selectedService, ok := s.apiServices[service.Site]
	if !ok {
		return
	}
	fields := dto.RequiredFields{
		UserID:   s.userID,
		ClientID: s.requestHeader.ClientID,
	}
	log, err = selectedService.CheckLogin(fields)
	s.recordResult(&serviceResult, service.Code, log, err)
	chResult <- serviceResult
}

func (s *sm) CheckAuth() request_v2.ManagerResponse {
	chResult := make(chan request_v2.ServiceStats)
	var results []request_v2.ServiceStats

	authList := s.apiAuthService.GetClientAll(s.userID, s.requestHeader.ClientID)
	var sites []request_v2.SiteEntry
	for _, auth := range authList {
		sites = append(sites, request_v2.SiteEntry{Site: auth.Service})
	}
	for _, service := range sites {
		go s.asyncCheckAuth(service, chResult)
	}

	for range len(sites) {
		eachResult := <-chResult
		results = append(results, eachResult)
	}
	close(chResult)

	result := request_v2.ManagerResponse{
		ReqHeaderEntry: s.requestHeader,
		Results:        results,
	}
	result.SetOveralStatus()
	return result
}
