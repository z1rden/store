package kafka_service

import (
	"context"
	"loms/internal/loms/config"
	"loms/internal/loms/kafka_producer"
	"loms/internal/loms/repository/kafka_storage"
	"sync"
)

type Service interface {
	SendMessages(ctx context.Context)
	StopSendMessages() error
}

type service struct {
	kafkaStorage    kafka_storage.Storage
	kafkaProducer   kafka_producer.Producer
	sendMessagesWG  sync.WaitGroup
	sendMessageDone chan struct{}
	cfg             *config.Config // Если понадобится отправить какую-то мета-информацию или будет несколько топиков
}

func NewService(kafkaStorage kafka_storage.Storage, kafkaProducer kafka_producer.Producer, cfg *config.Config) Service {
	return &service{
		kafkaStorage:    kafkaStorage,
		kafkaProducer:   kafkaProducer,
		sendMessageDone: make(chan struct{}),
		cfg:             cfg,
	}
}
