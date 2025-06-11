package suite

import (
	"cart/internal/cart/client/loms_service"
	loms_service_mock "cart/internal/cart/client/loms_service/mocks"
	"cart/internal/cart/client/product_service"
	product_service_mock "cart/internal/cart/client/product_service/mocks"
	"cart/internal/cart/repository/cart_storage"
	cart_storage_mock "cart/internal/cart/repository/cart_storage/mocks"
	"cart/internal/cart/service/cart_service"
	cart_service_mock "cart/internal/cart/service/cart_service/mocks"
)

type suiteProvider struct {
	cartStorage        cart_storage.Storage
	cartStorageMock    *cart_storage_mock.StorageMock
	productService     product_service.Client
	productServiceMock *product_service_mock.ClientMock
	cartService        cart_service.Service
	cartServiceMock    *cart_service_mock.ServiceMock
	lomsService        loms_service.Client
	lomsServiceMock    *loms_service_mock.ClientMock
}

func NewSuiteProvider() *suiteProvider {
	return &suiteProvider{}
}

func (s *suiteProvider) GetCartStorageMock() *cart_storage_mock.StorageMock {
	if s.cartStorageMock == nil {
		s.cartStorageMock = &cart_storage_mock.StorageMock{}
	}

	return s.cartStorageMock
}

func (s *suiteProvider) GetCartStorage() cart_storage.Storage {
	if s.cartStorage == nil {
		s.cartStorage = s.GetCartStorageMock()
	}

	return s.cartStorage
}

func (s *suiteProvider) GetProductServiceMock() *product_service_mock.ClientMock {
	if s.productServiceMock == nil {
		s.productServiceMock = &product_service_mock.ClientMock{}
	}

	return s.productServiceMock
}

func (s *suiteProvider) GetProductService() product_service.Client {
	if s.productService == nil {
		s.productService = s.GetProductServiceMock()
	}

	return s.productService
}

func (s *suiteProvider) GetCartServiceMock() *cart_service_mock.ServiceMock {
	if s.cartServiceMock == nil {
		s.cartServiceMock = &cart_service_mock.ServiceMock{}
	}

	return s.cartServiceMock
}

func (s *suiteProvider) GetCartService() cart_service.Service {
	if s.cartService == nil {
		s.cartService = s.GetCartServiceMock()
	}

	return s.cartService
}

func (s *suiteProvider) GetLomsServiceMock() *loms_service_mock.ClientMock {
	if s.lomsServiceMock == nil {
		s.lomsServiceMock = &loms_service_mock.ClientMock{}
	}

	return s.lomsServiceMock
}

func (s *suiteProvider) GetLomsService() loms_service.Client {
	if s.lomsServiceMock == nil {
		s.lomsServiceMock = s.GetLomsServiceMock()
	}

	return s.lomsServiceMock
}
