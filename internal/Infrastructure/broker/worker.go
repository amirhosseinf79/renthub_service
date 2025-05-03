package broker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/hibiken/asynq"
)

type serverS struct {
	client         interfaces.BrokerClientInterface
	server         *asynq.Server
	serviceManager interfaces.ServiceManager
	logger         interfaces.LoggerInterface
}

func (s *serverS) updateHandler(ctx context.Context, t *asynq.Task) error {
	fmt.Println(t.Type())
	var p dto.ClientUpdateBody
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("%v: %w", err, asynq.SkipRetry)
	}
	var result dto.ManagerResponse
	serviceManager := s.serviceManager.SetConfigs(p.UserID, p.Header, p.Services, p.Dates)

	switch t.Type() {
	case "update:calendar":
		result = serviceManager.CalendarUpdate(p.Action)
	case "update:minNight":
		result = serviceManager.MinNightUpdate(p.LimitDays)
	case "update:discount":
		result = serviceManager.DiscountUpdate(p.DiscountPercent)
	case "update:price":
		result = serviceManager.PriceUpdate()
	case "update:checkAuth":
		result = serviceManager.CheckAuth()
	default:
		return fmt.Errorf("unexpected task type: %s %w", t.Type(), asynq.SkipRetry)
	}
	p.FinalResult = result
	if len(result.Results) == 0 {
		return nil
	}
	fmt.Println("Done")
	s.client.AsyncUpdate("webhook", p)
	return nil
}

func (s *serverS) sendWebhook(ctx context.Context, t *asynq.Task) error {
	fmt.Println("update:webhook")
	var p dto.ClientUpdateBody
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("%v: %w", err, asynq.SkipRetry)
	}
	serviceManager := s.serviceManager.SetConfigs(p.UserID, p.Header, p.Services, p.Dates)
	log, err := serviceManager.SendWebhook(p.FinalResult)
	s.logger.RecordLog(log)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func NewWorker(
	client interfaces.BrokerClientInterface,
	sericeManager interfaces.ServiceManager,
	logger interfaces.LoggerInterface,
) interfaces.BrokerServerInterface {
	redisServer := os.Getenv("RedisServer")
	redisPass := os.Getenv("RedisPass")
	return &serverS{
		serviceManager: sericeManager,
		client:         client,
		logger:         logger,
		server: asynq.NewServer(
			asynq.RedisClientOpt{Addr: redisServer, Password: redisPass},
			asynq.Config{
				Concurrency: 10,
				RetryDelayFunc: func(n int, e error, t *asynq.Task) time.Duration {
					return 5 * time.Second
				},
			},
		),
	}
}

func (s *serverS) StartWorker() {
	mux := asynq.NewServeMux()
	mux.HandleFunc("update:webhook", s.sendWebhook)
	mux.HandleFunc("update:", s.updateHandler)

	if err := s.server.Run(mux); err != nil {
		log.Fatal(err)
	}
}
