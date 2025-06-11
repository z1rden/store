package service_provider

import "cart/internal/cart/repository/cart_storage"

type repository struct {
	cartStorage cart_storage.Storage
}

func (s *ServiceProvider) GetCartStorage() cart_storage.Storage {
	if s.repository.cartStorage == nil {
		s.repository.cartStorage = cart_storage.NewStorage()
	}

	return s.repository.cartStorage
}
