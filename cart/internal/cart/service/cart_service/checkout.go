package cart_service

import (
	"cart/internal/cart/client/loms_service"
	"cart/internal/cart/logger"
	"context"
	"fmt"
)

func (s *service) Checkout(ctx context.Context, userID int64) (int64, error) {
	const operation = "cart_service.checkout"

	cart, err := s.cartStorage.GetCartByUserID(ctx, userID)
	if err != nil {
		logger.Errorf(ctx, "%s: failed to get cart items by userID: %v", operation, err)
		return 0, fmt.Errorf("failed to get cart items by userID: %w", err)
	}

	orderID, err := s.lomsService.OrderCreate(ctx, userID, ToOrderItems(cart.Items))
	if err != nil {
		logger.Errorf(ctx, "%s: failed to create order: %v", operation, err)
		return 0, fmt.Errorf("failed to create order: %w", err)
	}

	err = s.cartStorage.DeleteCartByUserID(ctx, userID)
	if err != nil {
		logger.Errorf(ctx, "%s: failed to delete cart items by userID: %v", operation, err)
		return 0, fmt.Errorf("failed to delete cart items by userID: %w", err)
	}

	return orderID, nil
}

func ToOrderItems(items map[int64]uint16) []*loms_service.OrderItem {
	res := make([]*loms_service.OrderItem, 0, len(items))
	for sku, quantity := range items {
		res = append(res, &loms_service.OrderItem{
			Sku:      sku,
			Quantity: quantity,
		})
	}

	return res
}
