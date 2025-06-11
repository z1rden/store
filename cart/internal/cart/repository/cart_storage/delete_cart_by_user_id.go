package cart_storage

import (
	"context"
)

func (s *storage) DeleteCartByUserID(ctx context.Context, userID int64) error {
	s.Lock()
	defer s.Unlock()

	_, exists := s.cartStorage[userID]
	if exists {
		delete(s.cartStorage, userID)
	}

	return nil
}
