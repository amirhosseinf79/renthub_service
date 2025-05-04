package interfaces

import "github.com/amirhosseinf79/renthub_service/internal/dto"

type BrokerClientInterface interface {
	AsyncUpdate(task string, body dto.ClientUpdateBody)
}

type BrokerServerInterface interface {
	StartWorker()
}
