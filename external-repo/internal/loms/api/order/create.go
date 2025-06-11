package order_api

import (
	"buf.build/go/protovalidate"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"loms/pkg/api/order"
)

func (a *api) Create(ctx context.Context, r *order.OrderCreateRequest) (*order.OrderCreateResponse, error) {
	if err := protovalidate.Validate(r); err != nil {
		return nil, err
	}

	orderID, err := a.orderService.Create(ctx, r.User, ToOrderServiceItems(r.GetItems()))
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	return &order.OrderCreateResponse{
		OrderId: orderID,
	}, nil
}
