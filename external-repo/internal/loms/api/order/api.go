package order_api

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"loms/internal/loms/service/order_service"
	"loms/pkg/api/order"
)

type API interface {
	RegisterGrpcServer(s *grpc.Server)
	RegisterHttpHandlers(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
	Create(ctx context.Context, r *order.OrderCreateRequest) (*order.OrderCreateResponse, error)
	Info(ctx context.Context, r *order.OrderInfoRequest) (*order.OrderInfoResponse, error)
	Cancel(ctx context.Context, r *order.OrderCancelRequest) (*emptypb.Empty, error)
	Pay(ctx context.Context, r *order.OrderPayRequest) (*emptypb.Empty, error)
}

type api struct {
	order.UnimplementedOrderServer
	orderService order_service.Service
}

func NewApi(ctx context.Context, orderService order_service.Service) API {
	return &api{
		orderService: orderService,
	}
}

func (a *api) RegisterGrpcServer(s *grpc.Server) {
	order.RegisterOrderServer(s, a)
}

func (a *api) RegisterHttpHandlers(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	// RegisterHandler связывает соединение сервера gRPC, которое проксируется через gRPC-Gateway.  1
	// RegisterHandlerClient, в свою очередь, связывает клиент gRPC, и gRPC-Gateway преобразует HTTP-запросы в вызовы к этому клиенту.  1
	// Таким образом, основное отличие заключается в том, что первый метод работает с подключением сервера, а второй — с подключением клиента.
	err := order.RegisterOrderHandler(ctx, mux, conn)
	if err != nil {
		return err
	}

	return nil
}
