package loms_service

import (
	"cart/pkg/api/order"
	"cart/pkg/api/stock"
	"context"
	"google.golang.org/grpc"
)

type Client interface {
	OrderCreate(ctx context.Context, user int64, items []*OrderItem) (int64, error)
	StockInfo(ctx context.Context, SkuID int64) (uint16, error)
}

type client struct {
	orderGrpcClient order.OrderClient
	stockGrpcClient stock.StockClient
}

func NewClient(conn *grpc.ClientConn) Client {
	return &client{
		orderGrpcClient: order.NewOrderClient(conn),
		stockGrpcClient: stock.NewStockClient(conn),
	}
}
