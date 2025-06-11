package loms_service

import (
	"context"
)

func (c *client) OrderCreate(ctx context.Context, user int64, items []*OrderItem) (int64, error) {
	req := ToOrderCreateRequest(user, items)

	resp, err := c.orderGrpcClient.Create(ctx, req)
	if err != nil {
		return 0, err
	}

	return resp.GetOrderId(), nil
}
