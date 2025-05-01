package manager

import (
	"fmt"
	"time"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/internal/services/requests"
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

func (s *sm) sendWebhook(response dto.ManagerResponse) (log *models.Log, err error) {
	log = &models.Log{
		UserID:   s.userID,
		ClientID: response.ClientID,
		Service:  "webhook",
	}

	header := map[string]string{}
	extraH := map[string]string{}
	request := requests.New("POST", response.CallbackUrl, header, extraH, log)
	err = request.Start(response, "body")
	if err != nil {
		log.FinalResult = err.Error()
		return log, err
	}
	ok, err := request.Ok()
	if !ok {
		log.FinalResult = err.Error()
		return log, err
	}
	log.FinalResult = "success"
	log.IsSucceed = true
	return log, nil
}

func (s *sm) tryWebHook(result dto.ManagerResponse) {
	for try := range 5 {
		fmt.Println("Try:", try+1)
		log, err := s.sendWebhook(result)
		s.logger.RecordLog(log)
		if err != nil {
			// fmt.Printf("%v\nTrying in 5s...\n", err)
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}
}
