package manager

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (s *sm) asyncCalendar(service dto.SiteEntry, action string) (serviceResult dto.ServiceStats) {
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
		err = dto.ErrInvalidRequest
	}
	s.recordResult(&serviceResult, service.Code, log, err)
	return
}

func (s *sm) CalendarUpdate(action string) dto.ManagerResponse {
	var results []dto.ServiceStats
	for _, service := range s.services {
		// go s.asyncCalendar(service, action)
		result := s.asyncCalendar(service, action)
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
