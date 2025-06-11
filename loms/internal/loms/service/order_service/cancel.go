package order_service

import (
	"context"
	"loms/internal/loms/model"
)

func (s *service) Cancel(ctx context.Context, orderID int64) error {
	orderStorage, err := s.orderStorage.GetByID(ctx, orderID)
	if err != nil {
		return err
	}

	mOrder := ToModelOrder(orderStorage)

	if err := s.stockStorage.ReserveCancel(ctx, ToStockStorageItems(mOrder.Items)); err != nil {
		return err
	}

	if err := s.orderStorage.SetStatus(ctx, orderStorage.OrderID, model.OrderStatusCanceled); err != nil {
		return err
	}

	return nil
}
