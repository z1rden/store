package stock_storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"loms/internal/loms/model"
	"loms/internal/loms/repository/stock_storage/sqlc"
)

func (s *storage) Reserve(ctx context.Context, items []*ReserveItem) error {
	pool := s.dbClient.GetWriterPool()
	err := pool.BeginFunc(ctx, func(tx pgx.Tx) error {
		queries := sqlc.New(pool).WithTx(tx)

		for _, item := range items {
			storageItem, err := queries.GetBySku(ctx, item.SkuID)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					return model.ErrNotFound
				}

				return err
			}

			if storageItem.TotalCount-storageItem.Reserved < int32(item.Quantity) {
				return errors.New("reservation limit reached")
			}

			err = queries.Reserve(ctx, sqlc.ReserveParams{
				SkuID:    item.SkuID,
				Reserved: int32(item.Quantity),
			})
			if err != nil {
				return fmt.Errorf("failed to reserve stock for sku %d: %w", item.SkuID, err)
			}
		}

		return nil
	})

	if err != nil {
		return errors.New("failed to reserve")
	}

	return nil
}
