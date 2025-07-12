package broker_v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	request_v1 "github.com/amirhosseinf79/renthub_service/internal/dto/request/v1"
	"github.com/hibiken/asynq"
)

type serverS struct {
	client         interfaces.BrokerClientInterface
	server         *asynq.Server
	serices        map[string]interfaces.ApiService
	serviceManager interfaces.ServiceManager
	logger         interfaces.LoggerInterface
	webhookService interfaces.WebhookService
}

func NewWorker(
	client interfaces.BrokerClientInterface,
	sericeManager interfaces.ServiceManager,
	logger interfaces.LoggerInterface,
	serices map[string]interfaces.ApiService,
	webhookService interfaces.WebhookService,
) interfaces.BrokerServerInterface {
	redisServer := os.Getenv("RedisServer")
	redisPass := os.Getenv("RedisPass")
	return &serverS{
		serices:        serices,
		serviceManager: sericeManager,
		client:         client,
		logger:         logger,
		webhookService: webhookService,
		server: asynq.NewServer(
			asynq.RedisClientOpt{
				Addr:     redisServer,
				Password: redisPass,
				DB:       0,
			},
			asynq.Config{
				Concurrency: 50,
				RetryDelayFunc: func(n int, e error, t *asynq.Task) time.Duration {
					if errors.Is(e, dto.ErrorUnauthorized) {
						return 1 * time.Second
					}

					return 5 * time.Second
				},
			},
		),
	}
}

func (s *serverS) StartWorker() {
	mux := asynq.NewServeMux()
	mux.HandleFunc("otp:send", s.otpSendHandler)
	mux.HandleFunc("otp:verify", s.otpVerifyHandler)
	mux.HandleFunc("update:webhook", s.sendWebhook)
	mux.HandleFunc("update:", s.updateHandler)

	if err := s.server.Run(mux); err != nil {
		log.Fatal(err)
	}
}

func (s *serverS) updateHandler(ctx context.Context, t *asynq.Task) error {
	fmt.Println(t.Type())
	var p request_v1.ClientUpdateBody
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("%v: %w", err, asynq.SkipRetry)
	}
	var result request_v1.ManagerResponse
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
	case "update:token":
		result = serviceManager.ManageAutoLogin()
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

func (s *serverS) otpSendHandler(ctx context.Context, t *asynq.Task) error {
	fmt.Println(t.Type())
	var p request_v1.OTPBody
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("%v: %w", err, asynq.SkipRetry)
	}
	selectedService, ok := s.serices[p.Service]
	if !ok {
		return nil
	}
	log, _ := selectedService.SendOtp(dto.RequiredFields{UserID: p.UserID, ClientID: p.ClientID}, p.PhoneNumebr)
	s.logger.RecordLog(log)
	return nil
}

func (s *serverS) otpVerifyHandler(ctx context.Context, t *asynq.Task) error {
	fmt.Println(t.Type())
	var p request_v1.OTPBody
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("%v: %w", err, asynq.SkipRetry)
	}
	selectedService, ok := s.serices[p.Service]
	if !ok {
		return nil
	}
	log, _ := selectedService.VerifyOtp(dto.RequiredFields{
		UserID:   p.UserID,
		ClientID: p.ClientID,
	},
		dto.OTPCreds{
			PhoneNumber: p.PhoneNumebr,
			OTPCode:     p.Code,
		})
	s.logger.RecordLog(log)
	return nil
}

func (s *serverS) sendWebhook(ctx context.Context, t *asynq.Task) error {
	fmt.Println("update:webhook")
	var p request_v1.ClientUpdateBody
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("%v: %w", err, asynq.SkipRetry)
	}
	log, err := s.webhookService.SendResult(p)
	s.logger.RecordLog(log)
	if err != nil {
		if log.StatusCode == 401 {
			log, err2 := s.webhookService.RefreshToken(p.UserID)
			s.logger.RecordLog(log)
			if err2 != nil {
				fmt.Println(err2)
			}
			return fmt.Errorf("err: %w %w", err, dto.ErrorUnauthorized)
		}
		fmt.Println(err)
		return err
	}
	return nil
}
