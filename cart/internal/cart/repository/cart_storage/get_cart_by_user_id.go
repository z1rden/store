package cart_storage

import (
	"cart/internal/cart/model"
	"context"
)

func (s *storage) GetCartByUserID(ctx context.Context, userID int64) (*Cart, error) {
	s.RLock()
	defer s.RUnlock()

	cart, exists := s.cartStorage[userID]
	if !exists {
		return nil, model.ErrNotFound
	}

	return cart, nil
}
