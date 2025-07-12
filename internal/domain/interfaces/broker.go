package interfaces

import (
	request_v1 "github.com/amirhosseinf79/renthub_service/internal/dto/request/v1"
	request_v2 "github.com/amirhosseinf79/renthub_service/internal/dto/request/v2"
)

type BrokerClientInterface interface {
	AsyncUpdate(task string, body request_v1.ClientUpdateBody)
	AsyncOTP(task string, body request_v1.OTPBody)
}

type BrokerClientInterface_v2 interface {
	AsyncUpdate(task string, body request_v2.ClientUpdateBody)
	AsyncOTP(task string, body request_v2.OTPBody)
}

type BrokerServerInterface interface {
	StartWorker()
}
