package manager

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (s *sm) asyncCheckAuth(service dto.SiteEntry) (serviceResult dto.ServiceStats) {
	serviceResult = s.initServiceStatus(service.Site)
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
	s.recordResult(&serviceResult, service.Code, log, err)
	return
}

func (s *sm) CheckAuth() dto.ManagerResponse {
	var results []dto.ServiceStats
	authList := s.apiAuthService.GetClientAll(s.userID, s.responseHead.ClientID)
	var sites []dto.SiteEntry
	for _, auth := range authList {
		sites = append(sites, dto.SiteEntry{Site: auth.Service})
	}
	for _, service := range sites {
		// go s.asyncCheckAuth(service)
		result := s.asyncCheckAuth(service)
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
