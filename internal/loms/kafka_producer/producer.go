package kafka_producer

import (
	"context"
	"github.com/IBM/sarama"
)

type Producer interface {
	SendMessageWithKey(ctx context.Context, topic, key string, message interface{}, appName string) error
	Close() error
}

type producer struct {
	syncProducer sarama.SyncProducer
}

func NewProducer(addrs []string, options ...ProducerOption) (Producer, error) {
	config := NewConfig(options...)

	syncProducer, err := sarama.NewSyncProducer(addrs, config)
	if err != nil {
		return nil, err
	}

	return &producer{
		syncProducer: syncProducer,
	}, nil
}
