package order_storage

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5"
	"loms/internal/loms/model"
	"loms/internal/loms/repository/kafka_storage"
)

func (s *storage) insertOrderStatusChangedKafkaOutbox(ctx context.Context, tx pgx.Tx, orderID int64, status string) error {
	message := &model.OrderChangeStatusMessageOrder{
		OrderID: orderID,
		Status:  status,
	}

	messageJson, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	messageKafka := &kafka_storage.Message{
		Event:      model.EventOrderStatusChanged,
		EntityType: model.EntityTypeOrder,
		EntityID:   fmt.Sprintf("%d", orderID),
		Data:       string(messageJson),
	}

	err = kafka_storage.InsertMessage(ctx, tx, messageKafka)
	if err != nil {
		return fmt.Errorf("failed to insert message: %w", err)
	}

	return nil
}
