package manager

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (s *sm) asyncCalendar(service dto.SiteEntry, action string, chResult chan dto.ServiceStats) {
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
		err = dto.ErrInvalidRequest
	}
	s.recordResult(&serviceResult, service.Code, log, err)
	chResult <- serviceResult
}

func (s *sm) CalendarUpdate(action string) dto.ManagerResponse {
	chResult := make(chan dto.ServiceStats)
	var results []dto.ServiceStats

	for _, service := range s.services {
		go s.asyncCalendar(service, action, chResult)
	}

	for range len(s.services) {
		results = append(results, <-chResult)
	}
	close(chResult)

	result := dto.ManagerResponse{
		ReqHeaderEntry: s.responseHead,
		Results:        results,
	}
	result.SetOveralStatus()
	s.tryWebHook(result)
	return result
}
