package manager

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (s *sm) MinNightUpdate(limitDays int) dto.ManagerResponse {
	var results []dto.ServiceStats
	for _, service := range s.services {
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
			continue
		}

		switch limitDays {
		case 0:
			log, err = selectedService.UnsetMiniNight(fields)
		default:
			log, err = selectedService.SetMinNight(fields)
		}

		serviceResult.Code = service.Code
		s.logger.RecordLog(log)
		if err != nil {
			serviceResult.Status = "failed"
			serviceResult.ErrorMessage = err.Error()
		}
		results = append(results, serviceResult)
	}

	result := dto.ManagerResponse{
		ReqHeaderEntry: s.responseHead,
		Results:        results,
	}
	result.SetOveralStatus()
	return result
}
