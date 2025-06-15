package service_provider

import (
	"context"
	"github.com/IBM/sarama"
	"loms/internal/loms/kafka_admin"
	"loms/internal/loms/kafka_producer"
	"loms/internal/loms/logger"
	"time"
)

type kafka struct {
	admin    kafka_admin.Admin
	producer kafka_producer.Producer
}

func (s *ServiceProvider) GetKafkaProducer(ctx context.Context) kafka_producer.Producer {
	if s.kafka.producer == nil {
		var err error
		s.kafka.producer, err = kafka_producer.NewProducer(
			[]string{s.cfg.Kafka.Addr},
			kafka_producer.WithIdempotent(),
			kafka_producer.WithRequiredAcks(sarama.RequiredAcks(s.cfg.Kafka.Producer.RequiredAcks)),
			kafka_producer.WithMaxOpenRequests(s.cfg.Kafka.Producer.MaxOpenRequests),
			kafka_producer.WithMaxRetries(s.cfg.Kafka.Producer.MaxRetries),
			kafka_producer.WithRetryBackoff(time.Duration(s.cfg.Kafka.Producer.RetryBackoff)*time.Millisecond),
		)
		if err != nil {
			logger.Fatalf(ctx, "failed to create kafka producer: %v", err)
		}

		s.GetCloser(ctx).Add(s.kafka.producer.Close)
	}

	return s.kafka.producer
}

func (s *ServiceProvider) GetKafkaAdmin(ctx context.Context) kafka_admin.Admin {
	if s.kafka.admin == nil {
		var err error
		s.kafka.admin, err = kafka_admin.NewAdmin(
			[]string{s.cfg.Kafka.Addr},
		)

		if err != nil {
			logger.Fatalf(ctx, "failed to create kafka admin: %v", err)
		}
	}

	return s.kafka.admin
}
