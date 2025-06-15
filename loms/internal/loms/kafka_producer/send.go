package kafka_producer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"time"
)

func (p *producer) SendMessageWithKey(ctx context.Context, topic, key string, message interface{}, appName string) error {
	msg, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("encode message to json: %w", err)
	}

	_, _, err = p.syncProducer.SendMessage(&sarama.ProducerMessage{
		Headers: []sarama.RecordHeader{
			{
				Key:   sarama.ByteEncoder("app-name"),
				Value: sarama.ByteEncoder(appName),
			},
		},
		Key:       sarama.StringEncoder(key),
		Value:     sarama.ByteEncoder(msg),
		Topic:     topic,
		Timestamp: time.Now(),
	})

	if err != nil {
		return fmt.Errorf("send message to kafka failed: %w", err)
	}

	return nil
}

func (p *producer) Close() error {
	err := p.syncProducer.Close()
	if err != nil {
		return fmt.Errorf("close kafka producer failed: %w", err)
	}

	return nil
}
