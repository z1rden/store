package order_service

import (
	"context"
	"loms/internal/loms/model"
)

func (s *service) Info(ctx context.Context, orderID int64) (*model.Order, error) {
	order, err := s.orderStorage.GetByID(ctx, orderID)
	if err != nil {
		return nil, err
	}

	return ToModelOrder(order), nil
}
