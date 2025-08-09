package recieve_manager_v2

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	receive_manager_dto "github.com/amirhosseinf79/renthub_service/internal/dto/receive_manager"
	request_v2 "github.com/amirhosseinf79/renthub_service/internal/dto/request/v2"
)

type sm struct {
	apiAuthService interfaces.ApiAuthInterface
	logger         interfaces.LoggerInterface
	apiServices    map[string]interfaces.ApiService
	requestHeader  request_v2.ReqHeaderEntry
	timeLimit      int64
	userID         uint
	services       []request_v2.SiteRecieve
}

func New(
	apiServices map[string]interfaces.ApiService,
	apiAuthService interfaces.ApiAuthInterface,
	logger interfaces.LoggerInterface,
) interfaces.ServiceRecieveManagerV2 {
	return &sm{
		apiAuthService: apiAuthService,
		apiServices:    apiServices,
		logger:         logger,
		timeLimit:      15,
	}
}

func (s *sm) SetConfigs(
	userID uint,
	header request_v2.ReqHeaderEntry,
	services []request_v2.SiteRecieve,
) interfaces.ServiceRecieveManagerV2 {
	return &sm{
		apiAuthService: s.apiAuthService,
		apiServices:    s.apiServices,
		logger:         s.logger,
		timeLimit:      s.timeLimit,
		requestHeader:  header,
		services:       services,
		userID:         userID,
	}
}

func (s *sm) recordResult(log *models.Log, err error, response any) receive_manager_dto.SiteResponse {
	log.UpdateID = s.requestHeader.UpdateId
	s.logger.RecordLog(log)
	if err != nil {
		return receive_manager_dto.SiteResponse{
			ResponseCode: log.StatusCode,
			ErrorMessage: err.Error(),
		}
	}
	return receive_manager_dto.SiteResponse{
		Success:      true,
		ResponseCode: log.StatusCode,
		Response:     response,
	}
}
