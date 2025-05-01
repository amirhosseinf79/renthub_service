package manager

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (s *sm) asyncMinNight(service dto.SiteEntry, limitDays int, chResult chan dto.ServiceStats) {
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
		Amount: limitDays,
	}

	selectedService, ok := s.apiServices[service.Site]
	if !ok {
		return
	}

	switch limitDays {
	case 0:
		log, err = selectedService.UnsetMiniNight(fields)
	default:
		log, err = selectedService.SetMinNight(fields)
	}
	s.recordResult(&serviceResult, service.Code, log, err)
	chResult <- serviceResult
}

func (s *sm) MinNightUpdate(limitDays int) dto.ManagerResponse {
	chResult := make(chan dto.ServiceStats)
	var results []dto.ServiceStats

	for _, service := range s.services {
		go s.asyncMinNight(service, limitDays, chResult)

	}

	for range len(s.services) {
		results = append(results, <-chResult)
	}
	close(chResult)

	result := dto.ManagerResponse{
		ReqHeaderEntry: s.responseHead,
		OveralStatus:   "operating",
		Results:        results,
	}
	result.SetOveralStatus()
	s.tryWebHook(result)
	return result
}
