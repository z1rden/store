package order_storage

import (
	"context"
	"loms/internal/loms/db"
)

type Storage interface {
	Create(ctx context.Context, orderID int64, items []*Item) (int64, error)
	SetStatus(ctx context.Context, orderID int64, status string) error
	GetByID(ctx context.Context, orderID int64) (*Order, error)
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
