package broker_v2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	receive_manager_dto "github.com/amirhosseinf79/renthub_service/internal/dto/receive_manager"
	request_v2 "github.com/amirhosseinf79/renthub_service/internal/dto/request/v2"
	"github.com/hibiken/asynq"
)

type serverS struct {
	client                interfaces.BrokerClientInterface_v2
	server                *asynq.Server
	serices               map[string]interfaces.ApiService
	serviceUpdateManager  interfaces.ServiceUpdateManagerV2
	serviceRecieveManager interfaces.ServiceRecieveManagerV2
	logger                interfaces.LoggerInterface
	webhookService        interfaces.WebhookService_v2
}

func NewWorker(
	client interfaces.BrokerClientInterface_v2,
	serviceUpdateManager interfaces.ServiceUpdateManagerV2,
	serviceRecieveManager interfaces.ServiceRecieveManagerV2,
	logger interfaces.LoggerInterface,
	serices map[string]interfaces.ApiService,
	webhookService interfaces.WebhookService_v2,
) interfaces.BrokerServerInterface {
	redisServer := os.Getenv("RedisServer")
	redisPass := os.Getenv("RedisPass")
	return &serverS{
		serices:               serices,
		serviceUpdateManager:  serviceUpdateManager,
		serviceRecieveManager: serviceRecieveManager,
		client:                client,
		logger:                logger,
		webhookService:        webhookService,
		server: asynq.NewServer(
			asynq.RedisClientOpt{
				Addr:     redisServer,
				Password: redisPass,
				DB:       1,
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
	mux.HandleFunc("recieve:webhook", s.sendWebhook)
	mux.HandleFunc("update:", s.updateHandler)
	mux.HandleFunc("recieve:", s.recieveHandler)

	if err := s.server.Run(mux); err != nil {
		log.Fatal(err)
	}
}

func (s *serverS) updateHandler(ctx context.Context, t *asynq.Task) error {
	fmt.Println(t.Type())
	var p request_v2.ClientUpdateBody
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("%v: %w", err, asynq.SkipRetry)
	}
	var result request_v2.ManagerResponse
	serviceManager := s.serviceUpdateManager.SetConfigs(p.UserID, p.Header, p.Services, p.Dates)

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
	p.WebhookBody = result
	if len(result.Results) == 0 {
		return nil
	}
	fmt.Println("Done")
	s.client.AsyncUpdate("webhook", p)
	return nil
}

func (s *serverS) recieveHandler(ctx context.Context, t *asynq.Task) error {
	var p request_v2.ClientRecieveBody
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("%v: %w", err, asynq.SkipRetry)
	}

	recieveManager := s.serviceRecieveManager.SetConfigs(p.UserID, p.Header, p.Services)
	var response receive_manager_dto.RecieveResponse
	switch t.Type() {
	case "recieve:reservation":
		response = recieveManager.GetReservations()
	default:
		return fmt.Errorf("unexpected task type: %s %w", t.Type(), asynq.SkipRetry)
	}
	p.WebhookBody = response
	s.client.AsyncRecieve("webhook", p)
	return nil
}

func (s *serverS) otpSendHandler(ctx context.Context, t *asynq.Task) error {
	fmt.Println(t.Type())
	var p request_v2.OTPBody
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
	var p request_v2.OTPBody
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
	fmt.Println(t.Type())
	var log *models.Log
	var err error

	var p1 request_v2.ClientUpdateBody
	var p2 request_v2.ClientRecieveBody

	err1 := json.Unmarshal(t.Payload(), &p1)
	err2 := json.Unmarshal(t.Payload(), &p2)

	var whFields dto.WebhookFields

	if err1 == nil {
		whFields = dto.WebhookFields{
			UserID:      p1.UserID,
			UpdateId:    p1.Header.UpdateId,
			CallbackUrl: p1.Header.CallbackUrl,
			Body:        p1.WebhookBody,
		}
		log, err = s.webhookService.SendResult(whFields)

	} else if err2 == nil {
		whFields = dto.WebhookFields{
			UserID:      p2.UserID,
			UpdateId:    p2.Header.UpdateId,
			CallbackUrl: p2.Header.CallbackUrl,
			Body:        p2.WebhookBody,
		}
		log, err = s.webhookService.SendResult(whFields)

	} else {
		return fmt.Errorf("%v: %w", err, asynq.SkipRetry)
	}

	s.logger.RecordLog(log)
	if err != nil {
		if log.StatusCode == 401 {
			log, err2 := s.webhookService.RefreshToken(whFields.UserID)
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
