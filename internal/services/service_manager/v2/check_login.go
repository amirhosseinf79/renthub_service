package manager_v2

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	request_v2 "github.com/amirhosseinf79/renthub_service/internal/dto/request/v2"
)

func (s *sm) asyncCheckAuth(serviceName string, chResult chan request_v2.ServiceStats) {
	serviceResult := s.initServiceStatus(serviceName)
	var log *models.Log
	var err error

	selectedService, ok := s.apiServices[serviceName]
	if !ok {
		return
	}
	fields := dto.RequiredFields{
		UserID:   s.userID,
		ClientID: s.requestHeader.ClientID,
	}
	log, err = selectedService.CheckLogin(fields)
	s.recordResult(&serviceResult, "", log, err)
	chResult <- serviceResult
}

func (s *sm) CheckAuth() request_v2.ManagerResponse {
	chResult := make(chan request_v2.ServiceStats)
	var results []request_v2.ServiceStats

	for service := range s.apiServices {
		go s.asyncCheckAuth(service, chResult)
	}

	for range len(s.apiServices) {
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
