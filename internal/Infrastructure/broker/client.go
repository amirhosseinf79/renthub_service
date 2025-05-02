package broker

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/hibiken/asynq"
)

type client struct {
	client *asynq.Client
}

func NewClient(server, password string) interfaces.BrokerClientInterface {
	return &client{
		client: asynq.NewClient(asynq.RedisClientOpt{
			Addr:     server,
			Password: password,
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
