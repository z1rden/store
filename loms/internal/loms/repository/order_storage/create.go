package order_storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"loms/internal/loms/model"
	"loms/internal/loms/repository/order_storage/sqlc"
)

func (s *storage) Create(ctx context.Context, userID int64, items []*Item) (int64, error) {
	pool := s.dbClient.GetWriterPool()

	var orderID int64
	err := pool.BeginFunc(ctx, func(tx pgx.Tx) error {
		queries := sqlc.New(pool).WithTx(tx)

		var err error
		orderID, err = queries.CreateOrder(ctx, sqlc.CreateOrderParams{
			UserID: userID,
			Status: model.OrderStatusNew,
		})
		if err != nil {
			return fmt.Errorf("failed add row to order table: %w", err)
		}

		for _, item := range items {
			err = queries.AddOrderItem(ctx, sqlc.AddOrderItemParams{
				OrderID: orderID,
				SkuID:   item.SkuID,
				Quantity: pgtype.Int4{
					Int32: int32(item.Quantity),
					Valid: true,
				},
			})
			if err != nil {
				return fmt.Errorf("failed to add row to order_item table for sku %d: %w", item.SkuID, err)
			}
		}

		err = s.insertOrderStatusChangedKafkaOutbox(ctx, tx, orderID, model.OrderStatusNew)
		if err != nil {
			return fmt.Errorf("failed to insert to kafka outbox record: %w", err)
		}

		return nil
	})
	if err != nil {
		return 0, fmt.Errorf("failed to create order: %w", err)
	}

	return orderID, nil
}
