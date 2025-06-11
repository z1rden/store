package service_provider

import (
	conn_client "cart/internal/cart/client"
	"cart/internal/cart/client/loms_service"
	"cart/internal/cart/client/product_service"
)

type client struct {
	productService product_service.Client
	lomsService    loms_service.Client
}

func (s *ServiceProvider) GetProductService() product_service.Client {
	if s.client.productService == nil {
		s.client.productService = product_service.NewClient()
	}

	return s.client.productService
}

func (s *ServiceProvider) GetLomsService(port string) loms_service.Client {
	if s.client.lomsService == nil {
		s.client.lomsService = loms_service.NewClient(conn_client.GetClientConn(s.ctx, port))
	}

	return s.client.lomsService
}
