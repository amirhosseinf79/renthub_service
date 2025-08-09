package update_manager_v2

import (
	"errors"
	"time"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	request_v2 "github.com/amirhosseinf79/renthub_service/internal/dto/request/v2"
)

func (s *sm) asyncCalendar(service request_v2.SiteEntry, action string, chResult chan request_v2.ServiceStats) {
	serviceResult := s.initServiceStatus(service.Site)
	var log *models.Log
	var err error

	fields := dto.UpdateFields{
		RequiredFields: dto.RequiredFields{
			UserID:   s.userID,
			ClientID: service.ClientID,
		},
		RoomID: service.Code,
		Dates:  s.dates,
		Amount: 0,
	}
	selectedService, ok := s.apiServices[service.Site]
	if !ok {
		return
	}

	savedTime := time.Now().Unix()
	currentTime := savedTime
	for currentTime-savedTime < s.timeLimit {
		currentTime = time.Now().Unix()
		switch action {
		case "block":
			log, err = selectedService.CloseCalendar(fields)
		case "unblock":
			log, err = selectedService.OpenCalendar(fields)
		default:
			err = dto.ErrInvalidRequest
		}
		s.recordResult(&serviceResult, service.Code, log, err)
		if err != nil {
			if errors.Is(err, dto.ErrTimeOut) {
				continue
			}
		}
		break
	}
	chResult <- serviceResult
}

func (s *sm) CalendarUpdate(action string) request_v2.ManagerResponse {
	chResult := make(chan request_v2.ServiceStats)
	var results []request_v2.ServiceStats

	for _, service := range s.services {
		go s.asyncCalendar(service, action, chResult)
	}

	for range len(s.services) {
		results = append(results, <-chResult)
	}
	close(chResult)

	result := request_v2.ManagerResponse{
		ReqHeaderEntry: s.requestHeader,
		Results:        results,
	}
	result.SetOveralStatus()
	return result
}
