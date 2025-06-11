package service_provider

import (
	"context"
	order_api "loms/internal/loms/api/order"
	stock_api "loms/internal/loms/api/stock"
)

type api struct {
	orderAPI order_api.API
	stockAPI stock_api.API
}

func (s *ServiceProvider) GetOrderAPI(ctx context.Context) order_api.API {
	if s.api.orderAPI == nil {
		s.api.orderAPI = order_api.NewApi(ctx, s.GetOrderService(ctx))
	}

	return s.api.orderAPI
}

func (s *ServiceProvider) GetStockAPI(ctx context.Context) stock_api.API {
	if s.api.stockAPI == nil {
		s.api.stockAPI = stock_api.NewApi(ctx, s.GetStockService(ctx))
	}

	return s.api.stockAPI
}
