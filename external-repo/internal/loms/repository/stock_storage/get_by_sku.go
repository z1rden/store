package stock_storage

import (
	"context"
	"database/sql"
	"errors"
	"loms/internal/loms/model"
	"loms/internal/loms/repository/stock_storage/sqlc"
)

func (s *storage) GetBySku(ctx context.Context, SkuID int64) (uint16, error) {
	pool := s.dbClient.GetReaderPool()
	queries := sqlc.New(pool)

	storageItem, err := queries.GetBySku(ctx, SkuID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, model.ErrNotFound
		}

		return 0, err
	}

	return uint16(storageItem.TotalCount - storageItem.Reserved), nil

}
