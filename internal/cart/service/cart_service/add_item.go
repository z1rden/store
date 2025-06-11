package cart_service

import (
	"cart/internal/cart/logger"
	"context"
	"fmt"
)

func (s *service) AddItem(ctx context.Context, userID int64, skuID int64, count uint16) error {
	const operation = "cart_service.AddItem"

	_, err := s.productService.GetProduct(ctx, skuID)
	if err != nil {
		logger.Errorf(ctx, "%s: failed to get product: %v", operation, err)

		return err
	}

	quantity, err := s.lomsService.StockInfo(ctx, skuID)
	if err != nil {
		logger.Errorf(ctx, "%s: failed to stock info: %v", operation, err)

		return err
	}

	if quantity < count {
		logger.Errorf(ctx, "%s: not enough quantity for sku: %d", operation, skuID)

		return fmt.Errorf("not enough quantity for sku: %d", skuID)
	}

	err = s.cartStorage.AddItem(ctx, userID, skuID, count)
	if err != nil {
		logger.Errorf(ctx, "%s: failed to add item: %v", operation, err)

		return err
	}

	return nil
}
