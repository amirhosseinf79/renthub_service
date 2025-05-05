package broker

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/hibiken/asynq"
)

type client struct {
	client *asynq.Client
}

func NewClient() interfaces.BrokerClientInterface {
	redisServer := os.Getenv("RedisServer")
	redisPass := os.Getenv("RedisPass")
	return &client{
		client: asynq.NewClient(asynq.RedisClientOpt{
			Addr:     redisServer,
			Password: redisPass,
		}),
	}
}

func (c *client) AsyncUpdate(task string, body dto.ClientUpdateBody) {
	payload, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}
	t1 := asynq.NewTask(fmt.Sprintf("update:%v", task), payload, asynq.MaxRetry(3))
	_, err = c.client.Enqueue(t1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	log.Printf("[*] Successfully enqueued task %v", task)
}

func (c *client) AsyncOTP(task string, body dto.OTPBody) {
	payload, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}
	t1 := asynq.NewTask(fmt.Sprintf("otp:%v", task), payload)
	_, err = c.client.Enqueue(t1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	log.Printf("[*] Successfully enqueued task %v", task)
}
