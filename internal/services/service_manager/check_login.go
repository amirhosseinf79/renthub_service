package manager

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (s *sm) asyncCheckAuth(service dto.SiteEntry) {
	var results []dto.ServiceStats
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

func (s *sm) CheckAuth() dto.ManagerResponse {
	authList := s.apiAuthService.GetClientAll(s.userID, s.responseHead.ClientID)
	var sites []dto.SiteEntry
	for _, auth := range authList {
		sites = append(sites, dto.SiteEntry{Site: auth.Service})
	}
	for _, service := range sites {
		go s.asyncCheckAuth(service)
	}

	result := dto.ManagerResponse{
		ReqHeaderEntry: s.responseHead,
		OveralStatus:   "operating",
	}
	return result
}
