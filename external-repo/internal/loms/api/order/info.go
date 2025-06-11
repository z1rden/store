package order_api

import (
	"buf.build/go/protovalidate"
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"loms/internal/loms/model"
	"loms/pkg/api/order"
)

func (a *api) Info(ctx context.Context, r *order.OrderInfoRequest) (*order.OrderInfoResponse, error) {
	if err := protovalidate.Validate(r); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	mOrder, err := a.orderService.Info(ctx, r.OrderId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return ToOrderInfoResponse(mOrder), nil
}
