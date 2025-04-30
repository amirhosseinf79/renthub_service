package manager

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

type sm struct {
	logger       interfaces.LoggerInterface
	apiServices  map[string]interfaces.ApiService
	responseHead dto.ReqHeaderEntry
	userID       uint
	services     []dto.SiteEntry
	dates        []string
}

func New(apiServices map[string]interfaces.ApiService, logger interfaces.LoggerInterface) interfaces.ServiceManager {
	return &sm{
		apiServices: apiServices,
		logger:      logger,
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
	s.responseHead = header
	s.services = services
	s.userID = userID
	s.dates = dates
	return s
}
