package stock_storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"loms/internal/loms/model"
	"loms/internal/loms/repository/stock_storage/sqlc"
)

func (s *storage) ReserveCancel(ctx context.Context, items []*ReserveItem) error {
	pool := s.dbClient.GetWriterPool()
	err := pool.BeginFunc(ctx, func(tx pgx.Tx) error {
		queries := sqlc.New(pool).WithTx(tx)
		for _, item := range items {
			storageItem, err := queries.GetBySku(ctx, item.SkuID)
			if err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					return model.ErrNotFound
				}

				return err
			}

			if storageItem.TotalCount < int32(item.Quantity) {
				return fmt.Errorf("insufficient stock for product fro sku %d", item.SkuID)
			}

			if storageItem.Reserved < int32(item.Quantity) {
				return fmt.Errorf("insufficient reserve for product for sku %d", item.SkuID)
			}

			err = queries.ReserveCancel(ctx, sqlc.ReserveCancelParams{
				SkuID:    item.SkuID,
				Reserved: int32(item.Quantity),
			})
			if err != nil {
				return fmt.Errorf("failed to remove reserve for sku %d: %w", item.SkuID, err)
			}
		}

		return nil
	})
	if err != nil {
		return errors.New("failed to cancel reserve items")
	}

	return nil
}
