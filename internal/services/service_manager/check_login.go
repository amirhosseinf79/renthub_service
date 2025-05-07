package manager

import (
	"fmt"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (s *sm) asyncCheckAuth(service dto.SiteEntry, chResult chan dto.ServiceStats) {
	serviceResult := s.initServiceStatus(service.Site)
	var log *models.Log
	var err error

	selectedService, ok := s.apiServices[service.Site]
	if !ok {
		return
	}
	fields := dto.RequiredFields{
		UserID:   s.userID,
		ClientID: s.responseHead.ClientID,
	}
	log, err = selectedService.CheckLogin(fields)
	fmt.Println(service, err)
	s.recordResult(&serviceResult, service.Code, log, err)
	chResult <- serviceResult
}

func (s *sm) CheckAuth() dto.ManagerResponse {
	chResult := make(chan dto.ServiceStats)
	var results []dto.ServiceStats

	authList := s.apiAuthService.GetClientAll(s.userID, s.responseHead.ClientID)
	var sites []dto.SiteEntry
	for _, auth := range authList {
		sites = append(sites, dto.SiteEntry{Site: auth.Service})
	}
	for _, service := range sites {
		go s.asyncCheckAuth(service, chResult)
	}

	for range len(sites) {
		eachResult := <-chResult
		if eachResult.Status != "success" {
			results = append(results, eachResult)
		}
	}
	close(chResult)

	result := dto.ManagerResponse{
		ReqHeaderEntry: s.responseHead,
		Results:        results,
	}
	result.SetOveralStatus()
	return result
}
