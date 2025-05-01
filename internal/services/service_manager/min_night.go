package manager

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (s *sm) asyncMinNight(service dto.SiteEntry, limitDays int) (serviceResult dto.ServiceStats) {
	serviceResult = s.initServiceStatus(service.Site)
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
	return
}

func (s *sm) MinNightUpdate(limitDays int) dto.ManagerResponse {
	var results []dto.ServiceStats
	for _, service := range s.services {
		// go s.asyncMinNight(service, limitDays)
		result := s.asyncMinNight(service, limitDays)
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
