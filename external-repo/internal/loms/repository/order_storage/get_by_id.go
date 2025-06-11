package order_storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"loms/internal/loms/model"
	"loms/internal/loms/repository/order_storage/sqlc"
)

func (s *storage) GetByID(ctx context.Context, orderID int64) (*Order, error) {
	pool := s.dbClient.GetReaderPool()
	queries := sqlc.New(pool)

	order, err := queries.GetOrderByOrderID(ctx, orderID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model.ErrNotFound
		}

		return nil, fmt.Errorf("failed to select order %d: %w", orderID, err)
	}

	items, err := queries.GetOrderItemsByOrderID(ctx, orderID)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return nil, model.ErrNotFound
		}

		return nil, fmt.Errorf("failed to select order items from order %d: %w", orderID, err)
	}

	return &Order{
		OrderID: order.OrderID,
		UserID:  order.UserID,
		Items:   toOrderItems(items),
		Status:  string(order.Status),
	}, nil
}
