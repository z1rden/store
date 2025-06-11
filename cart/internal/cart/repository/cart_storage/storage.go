package cart_storage

import (
	"context"
	"sync"
)

type Storage interface {
	AddItem(ctx context.Context, userID int64, skuID int64, count uint16) error
	DeleteItem(ctx context.Context, userID int64, skuID int64) error
	GetCartByUserID(ctx context.Context, userID int64) (*Cart, error)
	DeleteCartByUserID(ctx context.Context, userID int64) error
}

type storage struct {
	sync.RWMutex
	cartStorage map[int64]*Cart
}

func NewStorage() Storage {
	return &storage{
		cartStorage: map[int64]*Cart{},
	}
}
