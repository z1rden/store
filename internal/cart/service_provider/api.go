package service_provider

import (
	"cart/internal/cart/api/cart_api"
)

type api struct {
	cartAPI cart_api.API
}

func (s *ServiceProvider) GetApi(port string) cart_api.API {
	if s.api.cartAPI == nil {
		s.api.cartAPI = cart_api.NewApi(
			serviceProvider.GetCartService(port))
	}

	return s.api.cartAPI
}
