package cart_service

import (
	"cart/internal/cart/logger"
	"context"
)

func (s *service) DeleteItem(ctx context.Context, userID int64, skuID int64) error {
	const operation = "cart_service.DeleteItem"

	err := s.cartStorage.DeleteItem(ctx, userID, skuID)
	if err != nil {
		logger.Errorf(ctx, "%s: failed to delete item from cart: %v", operation, err)

		return err
	}

	return nil
}
