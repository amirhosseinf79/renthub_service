package manager_v1

import (
	"sort"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	request_v1 "github.com/amirhosseinf79/renthub_service/internal/dto/request/v1"
)

type sm struct {
	apiAuthService interfaces.ApiAuthInterface
	logger         interfaces.LoggerInterface
	apiServices    map[string]interfaces.ApiService
	requestHeader  request_v1.ReqHeaderEntry
	userID         uint
	services       []request_v1.SiteEntry
	dates          []string
}

func New(
	apiServices map[string]interfaces.ApiService,
	apiAuthService interfaces.ApiAuthInterface,
	logger interfaces.LoggerInterface,
) interfaces.ServiceManager {
	return &sm{
		apiAuthService: apiAuthService,
		apiServices:    apiServices,
		logger:         logger,
	}
}

func (s *sm) SetConfigs(
	userID uint,
	header request_v1.ReqHeaderEntry,
	services []request_v1.SiteEntry,
	dates []string,
) interfaces.ServiceManager {
	sort.Strings(dates)
	return &sm{
		apiAuthService: s.apiAuthService,
		apiServices:    s.apiServices,
		logger:         s.logger,
		requestHeader:  header,
		services:       services,
		userID:         userID,
		dates:          dates,
	}
}

func (s *sm) initServiceStatus(service string) request_v1.ServiceStats {
	return request_v1.ServiceStats{
		Status: "success",
		Site:   service,
	}
}

func (s *sm) recordResult(serviceResult *request_v1.ServiceStats, objId string, log *models.Log, err error) {
	serviceResult.Code = objId
	s.logger.RecordLog(log)
	if err != nil {
		serviceResult.Status = "failed"
		serviceResult.ErrorMessage = err.Error()
	}
}
