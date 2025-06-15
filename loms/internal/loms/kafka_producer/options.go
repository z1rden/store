package kafka_producer

import (
	"github.com/IBM/sarama"
	"time"
)

type ProducerOption func(*sarama.Config)

func WithPartitioner(pfn sarama.PartitionerConstructor) ProducerOption {
	return func(config *sarama.Config) {
		config.Producer.Partitioner = pfn
	}
}

func WithRequiredAcks(acks sarama.RequiredAcks) ProducerOption {
	return func(config *sarama.Config) {
		config.Producer.RequiredAcks = acks
	}
}

func WithIdempotent() ProducerOption {
	return func(config *sarama.Config) {
		config.Producer.Idempotent = true
	}
}

func WithMaxRetries(n int) ProducerOption {
	return func(config *sarama.Config) {
		config.Producer.Retry.Max = n
	}
}

func WithRetryBackoff(d time.Duration) ProducerOption {
	return func(config *sarama.Config) {
		config.Producer.Retry.Backoff = d
	}
}

func WithMaxOpenRequests(n int) ProducerOption {
	return func(config *sarama.Config) {
		config.Net.MaxOpenRequests = n
	}
}

// WithFlushMessages: устанавливает максимальное количество сообщений в буфере перед отправкой в Kafka.
func WithFlushMessages(n int) ProducerOption {
	return func(config *sarama.Config) {
		config.Producer.Flush.Messages = n
	}
}

// WithFlushFrequency: устанавливает интервал принудительной отправки сообщений, даже если буфер не заполнен.
func WithFlushFrequency(d time.Duration) ProducerOption {
	return func(config *sarama.Config) {
		config.Producer.Flush.Frequency = d
	}
}
