package cart_service

import (
	"cart/internal/cart/logger"
	"context"
)

func (s *service) DeleteCartByUserId(ctx context.Context, userID int64) error {
	const operation = "cart_service.DeleteCartByUserId"

	err := s.cartStorage.DeleteCartByUserID(ctx, userID)
	if err != nil {
		logger.Errorf(ctx, "%s: failed to delete cart by user id %v", operation, err)

		return err
	}

	return nil
}
