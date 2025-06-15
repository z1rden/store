package order_storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"loms/internal/loms/repository/order_storage/sqlc"
)

func (s *storage) SetStatus(ctx context.Context, orderID int64, status string) error {
	pool := s.dbClient.GetWriterPool()
	err := pool.BeginFunc(ctx, func(tx pgx.Tx) error {
		queries := sqlc.New(pool).WithTx(tx)

		err := queries.UpdateStatusOrderByOrderID(ctx, sqlc.UpdateStatusOrderByOrderIDParams{
			OrderID: orderID,
			Status:  sqlc.OrderStatusType(status),
		})

		err = s.insertOrderStatusChangedKafkaOutbox(ctx, tx, orderID, status)
		if err != nil {
			return fmt.Errorf("failed to insert to kafka outbox record: %w", err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to update order status for %d: %w", orderID, err)
	}

	return nil
}
