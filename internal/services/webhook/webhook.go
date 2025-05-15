package webhook

import (
	"errors"
	"fmt"
	"net"
	"net/url"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/internal/services/requests"
	"github.com/hibiken/asynq"
)

type webhookS struct {
	userService interfaces.UserService
}

func NewWebhookService(userService interfaces.UserService) interfaces.WebhookService {
	return &webhookS{
		userService: userService,
	}
}

func (w *webhookS) SendResult(response dto.ClientUpdateBody) (log *models.Log, err error) {
	log = &models.Log{
		UserID:   response.UserID,
		ClientID: response.Header.ClientID,
		Service:  "webhook",
	}

	userM, err := w.userService.GetUserById(response.UserID)
	if err != nil {
		log.FinalResult = err.Error()
		return
	}

	_, err = url.ParseRequestURI(response.Header.CallbackUrl)
	if err != nil {
		log.FinalResult = "ignored/invalid url"
		return log, fmt.Errorf("err: %s %w", err.Error(), asynq.SkipRetry)
	}

	header := map[string]string{
		"Authorization": "Bearer %v",
	}
	extraH := map[string]string{
		"Authorization": userM.HookToken,
	}
	request := requests.New("POST", response.Header.CallbackUrl, header, extraH, log)
	err = request.Start(response.FinalResult, "body")
	if err != nil {
		var dnsErr *net.DNSError
		if has := errors.As(err, &dnsErr); has {
			log.FinalResult = "ignored/invalid url"
			return log, fmt.Errorf("err: %s %w", err.Error(), asynq.SkipRetry)
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

func (w *webhookS) RefreshToken(userID uint) (log *models.Log, err error) {
	log = &models.Log{
		UserID:  userID,
		Service: "webhook",
	}

	userM, err := w.userService.GetUserById(userID)
	if err != nil {
		log.FinalResult = err.Error()
		return
	}

	header := map[string]string{}
	extraH := map[string]string{}

	body := dto.WebhookRefreshBody{
		RefreshToken: userM.HookRefresh,
	}

	request := requests.New("POST", userM.RefreshURL, header, extraH, log)
	err = request.Start(body, "body")
	if err != nil {
		var dnsErr *net.DNSError
		if has := errors.As(err, &dnsErr); has {
			log.FinalResult = "ignored/invalid url"
			return log, fmt.Errorf("err: %s %w", err.Error(), asynq.SkipRetry)
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
