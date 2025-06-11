package order_service

import (
	"context"
	"loms/internal/loms/model"
)

func (s *service) Pay(ctx context.Context, orderID int64) error {
	orderStorage, err := s.orderStorage.GetByID(ctx, orderID)
	if err != nil {
		return err
	}

	mOrder := ToModelOrder(orderStorage)

	if err := s.stockStorage.ReserveRemove(ctx, ToStockStorageItems(mOrder.Items)); err != nil {
		return err
	}

	if err := s.orderStorage.SetStatus(ctx, orderStorage.OrderID, model.OrderStatusPayed); err != nil {
		return err
	}

	return nil
}
