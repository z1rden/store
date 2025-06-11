package cart_storage

import "context"

func (s *storage) AddItem(ctx context.Context, userID int64, skuID int64, count uint16) error {
	s.Lock()
	defer s.Unlock()

	cart, ok := s.cartStorage[userID]
	if !ok {
		cart = &Cart{
			Items: map[int64]uint16{},
		}
		s.cartStorage[userID] = cart
	}

	cart.Items[skuID] += count

	return nil
}
