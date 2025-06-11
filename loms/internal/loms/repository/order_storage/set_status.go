package order_storage

import (
	"context"
	"fmt"
	"loms/internal/loms/repository/order_storage/sqlc"
)

func (s *storage) SetStatus(ctx context.Context, orderID int64, status string) error {
	pool := s.dbClient.GetWriterPool()
	queries := sqlc.New(pool)

	err := queries.UpdateStatusOrderByOrderID(ctx, sqlc.UpdateStatusOrderByOrderIDParams{
		OrderID: orderID,
		Status:  sqlc.OrderStatusType(status),
	})
	if err != nil {
		return fmt.Errorf("failed to update order status for %d: %w", orderID, err)
	}

	return nil
}
