package service_provider

import "cart/internal/cart/service/cart_service"

type service struct {
	cartService cart_service.Service
}

func (s *ServiceProvider) GetCartService(port string) cart_service.Service {
	if s.service.cartService == nil {
		s.service.cartService = cart_service.NewService(
			s.GetCartStorage(),
			s.GetProductService(),
			s.GetLomsService(port),
		)
	}

	return s.service.cartService
}
