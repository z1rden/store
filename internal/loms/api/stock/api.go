package stock_api

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"loms/internal/loms/service/stock_service"
	"loms/pkg/api/stock"
)

type API interface {
	RegisterGrpcServer(s *grpc.Server)
	RegisterHttpHandlers(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
	Info(ctx context.Context, r *stock.StockInfoRequest) (*stock.StockInfoResponse, error)
}

type api struct {
	stock.UnimplementedStockServer
	stockService stock_service.Service
}

func NewApi(ctx context.Context, stockService stock_service.Service) API {
	return &api{
		stockService: stockService,
	}
}

func (a *api) RegisterGrpcServer(s *grpc.Server) {
	stock.RegisterStockServer(s, a)
}

func (a *api) RegisterHttpHandlers(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	err := stock.RegisterStockHandler(ctx, mux, conn)
	if err != nil {
		return err
	}

	return nil
}
