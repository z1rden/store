package order_service

import (
	"context"
	"loms/internal/loms/model"
	"loms/internal/loms/repository/order_storage"
	"loms/internal/loms/repository/stock_storage"
)

type Service interface {
	Create(ctx context.Context, userID int64, items []*model.Item) (int64, error)
	Info(ctx context.Context, orderID int64) (*model.Order, error)
	Cancel(ctx context.Context, orderID int64) error
	Pay(ctx context.Context, orderID int64) error
}

type service struct {
	orderStorage order_storage.Storage
	stockStorage stock_storage.Storage
}

func NewService(ctx context.Context, orderStorage order_storage.Storage, stockStorage stock_storage.Storage) Service {
	return &service{
		orderStorage: orderStorage,
		stockStorage: stockStorage,
	}
}
