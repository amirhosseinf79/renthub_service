package manager

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (s *sm) asyncCalendar(service dto.SiteEntry, action string) {
	var results []dto.ServiceStats
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
		Amount: 0,
	}
	selectedService, ok := s.apiServices[service.Site]
	if !ok {
		return
	}

	switch action {
	case "block":
		log, err = selectedService.CloseCalendar(fields)
	case "unblock":
		log, err = selectedService.OpenCalendar(fields)
	default:
		return
	}
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

func (s *sm) CalendarUpdate(action string) dto.ManagerResponse {
	for _, service := range s.services {
		go s.asyncCalendar(service, action)
	}

	result := dto.ManagerResponse{
		ReqHeaderEntry: s.responseHead,
		OveralStatus:   "operating",
	}
	return result
}
