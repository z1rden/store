package order_api

import (
	"buf.build/go/protovalidate"
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"loms/internal/loms/model"
	"loms/pkg/api/order"
)

func (a *api) Pay(ctx context.Context, r *order.OrderPayRequest) (*emptypb.Empty, error) {
	if err := protovalidate.Validate(r); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := a.orderService.Pay(ctx, r.GetOrderId()); err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
