package stock_service

import (
	"context"
)

func (s *service) Info(ctx context.Context, SkuID int64) (uint16, error) {
	quantity, err := s.stockStorage.GetBySku(ctx, SkuID)
	if err != nil {
		return 0, err
	}

	return quantity, nil
}
