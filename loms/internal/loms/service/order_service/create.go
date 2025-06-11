package order_service

import (
	"context"
	"loms/internal/loms/model"
)

func (s *service) Create(ctx context.Context, userID int64, items []*model.Item) (int64, error) {
	orderID, err := s.orderStorage.Create(ctx, userID, ToOrderStorageItems(items))
	if err != nil {
		return 0, err
	}

	if err := s.stockStorage.Reserve(ctx, ToStockStorageItems(items)); err != nil {
		if err := s.orderStorage.SetStatus(ctx, orderID, model.OrderStatusFailed); err != nil {
			return 0, err
		}

		return 0, err
	} else {
		if err := s.orderStorage.SetStatus(ctx, orderID, model.OrderStatusAwaitingPayment); err != nil {
			return 0, err
		}
	}

	return orderID, nil
}
