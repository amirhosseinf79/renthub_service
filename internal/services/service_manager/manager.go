package manager

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

type sm struct {
	apiAuthService interfaces.ApiAuthInterface
	logger         interfaces.LoggerInterface
	apiServices    map[string]interfaces.ApiService
	responseHead   dto.ReqHeaderEntry
	userID         uint
	services       []dto.SiteEntry
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

func (s *sm) initServiceStatus(service string) dto.ServiceStats {
	return dto.ServiceStats{
		Status: "success",
		Site:   service,
	}
}

func (s *sm) SetConfigs(
	userID uint,
	header dto.ReqHeaderEntry,
	services []dto.SiteEntry,
	dates []string,
) interfaces.ServiceManager {
	return &sm{
		apiAuthService: s.apiAuthService,
		apiServices:    s.apiServices,
		logger:         s.logger,
		responseHead:   header,
		services:       services,
		userID:         userID,
		dates:          dates,
	}
}

func (s *sm) recordResult(serviceResult *dto.ServiceStats, statusCode string, log *models.Log, err error) {
	serviceResult.Code = statusCode
	s.logger.RecordLog(log)
	if err != nil {
		serviceResult.Status = "failed"
		serviceResult.ErrorMessage = err.Error()
	}
}
