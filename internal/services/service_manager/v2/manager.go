package manager_v2

import (
	"sort"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	request_v2 "github.com/amirhosseinf79/renthub_service/internal/dto/request/v2"
)

type sm struct {
	apiAuthService interfaces.ApiAuthInterface
	logger         interfaces.LoggerInterface
	apiServices    map[string]interfaces.ApiService
	requestHeader  request_v2.ReqHeaderEntry
	timeLimit      int64
	userID         uint
	services       []request_v2.SiteEntry
	dates          []string
}

func New(
	apiServices map[string]interfaces.ApiService,
	apiAuthService interfaces.ApiAuthInterface,
	logger interfaces.LoggerInterface,
) interfaces.ServiceManager_v2 {
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
	services []request_v2.SiteEntry,
	dates []string,
) interfaces.ServiceManager_v2 {
	sort.Strings(dates)
	return &sm{
		apiAuthService: s.apiAuthService,
		apiServices:    s.apiServices,
		logger:         s.logger,
		timeLimit:      s.timeLimit,
		requestHeader:  header,
		services:       services,
		userID:         userID,
		dates:          dates,
	}
}

func (s *sm) initServiceStatus(service string) request_v2.ServiceStats {
	return request_v2.ServiceStats{
		Status: "success",
		Site:   service,
	}
}

func (s *sm) recordResult(serviceResult *request_v2.ServiceStats, objId string, log *models.Log, err error) {
	serviceResult.Code = objId
	if s.requestHeader.ClientID == "" {
		serviceResult.ClientID = log.ClientID
	}
	s.logger.RecordLog(log)
	if err != nil {
		serviceResult.Status = "failed"
		serviceResult.ErrorMessage = err.Error()
	} else {
		serviceResult.Status = "success"
		serviceResult.ErrorMessage = ""
	}
}
