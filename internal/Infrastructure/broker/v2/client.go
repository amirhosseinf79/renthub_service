package broker_v2

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	request_v2 "github.com/amirhosseinf79/renthub_service/internal/dto/request/v2"
	"github.com/hibiken/asynq"
)

type client struct {
	client *asynq.Client
}

func NewClient() interfaces.BrokerClientInterface_v2 {
	redisServer := os.Getenv("RedisServer")
	redisPass := os.Getenv("RedisPass")
	return &client{
		client: asynq.NewClient(asynq.RedisClientOpt{
			Addr:     redisServer,
			Password: redisPass,
			DB:       1,
		}),
	}
}

func (c *client) AsyncUpdate(task string, body request_v2.ClientUpdateBody) {
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

func (c *client) AsyncOTP(task string, body request_v2.OTPBody) {
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
