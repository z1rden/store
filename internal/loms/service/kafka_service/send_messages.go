package kafka_service

import (
	"context"
	"encoding/json"
	"fmt"
	"loms/internal/loms/logger"
	"loms/internal/loms/model"
	"loms/internal/loms/repository/kafka_storage/sqlc"
	"time"
)

func (s *service) SendMessages(ctx context.Context) {
	logger.Info(ctx, "kafka outbox sender is starting...")
	s.sendMessagesWG.Add(1)
	go func() {
		s.sendMessages(ctx)
	}()
}

func (s *service) StopSendMessages() error {
	close(s.sendMessageDone)
	s.sendMessagesWG.Wait()

	return nil
}

func (s *service) sendMessages(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-s.sendMessageDone:
			s.sendMessagesWG.Done()
			logger.Info(ctx, "kafka outbox sender is stopped successfully")

			return
		case <-ticker.C:
			err := s.kafkaStorage.SendMessages(ctx, s.sendMessage)

			if err != nil {
				logger.Errorf(ctx, "failed to send messages: %v", err)
			}
		}
	}
}

func (s *service) sendMessage(ctx context.Context, message *sqlc.KafkaOutbox) error {
	var err error
	switch message.Event.String {
	case model.EventOrderStatusChanged:
		err = s.sendOrderStatusChangedMessage(ctx, message)
	default:
		logger.Errorf(ctx, "failed to send message: %v", err)
	}

	return err
}

func (s *service) sendOrderStatusChangedMessage(ctx context.Context, message *sqlc.KafkaOutbox) error {
	order := &model.OrderChangeStatusMessageOrder{}
	err := json.Unmarshal([]byte(message.Data.String), order)
	if err != nil {
		return fmt.Errorf("failed to unmarshal order data: %w", err)
	}

	orderStatusChangedMessage := model.OrderChangeStatusMessage{
		ID:         message.MessageID,
		Time:       message.CreatedAt.Time,
		Event:      message.Event.String,
		EntityType: message.EntityType.String,
		EntityID:   message.EntityID.String,
		Data: model.OrderChangeStatusMessageData{
			Order: *order,
		},
	}

	err = s.kafkaProducer.SendMessageWithKey(
		ctx,
		s.cfg.Kafka.Topic.Name,
		message.EntityID.String,
		orderStatusChangedMessage,
		s.cfg.AppName)

	if err != nil {
		logger.Errorf(ctx, "failed to send OrderChangeStatusMessage: %v", err)
		return fmt.Errorf("failed to send OrderChangeStatusMessage: %w", err)
	}

	return nil
}
