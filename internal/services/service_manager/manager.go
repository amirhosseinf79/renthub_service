package manager

import (
	"errors"
	"net"
	"net/url"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/internal/services/requests"
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

func (s *sm) SendWebhook(response dto.ManagerResponse) (log *models.Log, err error) {
	log = &models.Log{
		UserID:   s.userID,
		ClientID: response.ClientID,
		Service:  "webhook",
	}

	_, err = url.ParseRequestURI(response.CallbackUrl)
	if err != nil {
		log.FinalResult = "ignored/invalid url"
		return log, nil
	}

	header := map[string]string{}
	extraH := map[string]string{}
	request := requests.New("POST", response.CallbackUrl, header, extraH, log)
	err = request.Start(response, "body")
	if err != nil {
		var dnsErr *net.DNSError
		if has := errors.As(err, &dnsErr); has {
			log.FinalResult = "ignored/invalid url"
			return log, nil
		}
		log.FinalResult = err.Error()
		return
	}
	ok, err := request.Ok()
	if !ok {
		log.FinalResult = err.Error()
		return
	}
	log.FinalResult = "success"
	log.IsSucceed = true
	return log, nil
}

func (s *sm) recordResult(serviceResult *dto.ServiceStats, statusCode string, log *models.Log, err error) {
	serviceResult.Code = statusCode
	s.logger.RecordLog(log)
	if err != nil {
		serviceResult.Status = "failed"
		serviceResult.ErrorMessage = err.Error()
	}
}
