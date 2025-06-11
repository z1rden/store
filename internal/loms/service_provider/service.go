package service_provider

import (
	"context"
	"loms/internal/loms/service/order_service"
	"loms/internal/loms/service/stock_service"
)

type service struct {
	orderService order_service.Service
	stockService stock_service.Service
}

func (s *ServiceProvider) GetOrderService(ctx context.Context) order_service.Service {
	if s.service.orderService == nil {
		s.service.orderService = order_service.NewService(ctx, s.GetOrderStorage(ctx), s.GetStockStorage(ctx))
	}

	return s.service.orderService
}

func (s *ServiceProvider) GetStockService(ctx context.Context) stock_service.Service {
	if s.service.stockService == nil {
		s.service.stockService = stock_service.NewService(ctx, s.GetStockStorage(ctx))
	}

	return s.service.stockService
}
