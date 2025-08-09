package recieve_manager_v2

import (
	"errors"
	"time"

	"github.com/amirhosseinf79/renthub_service/internal/dto"
	receive_manager_dto "github.com/amirhosseinf79/renthub_service/internal/dto/receive_manager"
	request_v2 "github.com/amirhosseinf79/renthub_service/internal/dto/request/v2"
)

func (s *sm) asyncGetReservations(field request_v2.SiteRecieve, result chan receive_manager_dto.SiteResponse) {
	service, ok := s.apiServices[field.Site]
	if !ok {
		return
	}

	if field.Filters == nil {
		field.Filters = make(map[string]any)
	}

	fields := dto.RecieveFields{
		RequiredFields: dto.RequiredFields{
			UserID:   s.userID,
			ClientID: field.ClientID,
		},
		Filters: field.Filters,
	}
	var finalResponse receive_manager_dto.SiteResponse
	savedTime := time.Now().Unix()
	currentTime := savedTime
	for currentTime-savedTime < s.timeLimit {
		currentTime = time.Now().Unix()
		var response any
		log, err := service.GetReservations(fields, &response)
		finalResponse = s.recordResult(log, err, response)
		if err != nil {
			if errors.Is(err, dto.ErrTimeOut) {
				continue
			}
		}
		break
	}
	result <- finalResponse
}

func (s *sm) GetReservations() (response receive_manager_dto.RecieveResponse) {
	response = make(receive_manager_dto.RecieveResponse)
	siteResponse := make(chan receive_manager_dto.SiteResponse)
	defer close(siteResponse)

	for _, f := range s.services {
		go s.asyncGetReservations(f, siteResponse)
	}

	for _, service := range s.services {
		response[service.Site] = <-siteResponse
	}
	return
}
