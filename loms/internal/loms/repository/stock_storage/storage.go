package stock_storage

import (
	"context"
	"loms/internal/loms/db"
)

type Storage interface {
	Reserve(ctx context.Context, items []*ReserveItem) error
	GetBySku(ctx context.Context, SkuID int64) (uint16, error)
	ReserveCancel(ctx context.Context, items []*ReserveItem) error
	ReserveRemove(ctx context.Context, items []*ReserveItem) error
}

type storage struct {
	ctx      context.Context
	dbClient db.Client
}

func NewStorage(ctx context.Context, dbClient db.Client) Storage {
	return &storage{
		ctx:      ctx,
		dbClient: dbClient,
	}
}
