package stock_service

import (
	"context"
	"loms/internal/loms/repository/stock_storage"
)

type Service interface {
	Info(ctx context.Context, SkuID int64) (uint16, error)
}

type service struct {
	stockStorage stock_storage.Storage
}

func NewService(ctx context.Context, stockStorage stock_storage.Storage) Service {
	return &service{
		stockStorage: stockStorage,
	}
}
