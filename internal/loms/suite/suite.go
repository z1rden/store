package suite

import (
	"loms/internal/loms/repository/order_storage"
	order_storage_mock "loms/internal/loms/repository/order_storage/mocks"
	"loms/internal/loms/repository/stock_storage"
	stock_storage_mock "loms/internal/loms/repository/stock_storage/mocks"
)

type suiteProvider struct {
	orderStorageMock *order_storage_mock.StorageMock
	orderStorage     order_storage.Storage
	stockStorageMock *stock_storage_mock.StorageMock
	stockStorage     stock_storage.Storage
}

func NewSuiteProvider() *suiteProvider {
	return &suiteProvider{}
}

func (s *suiteProvider) GetOrderStorageMock() *order_storage_mock.StorageMock {
	if s.orderStorageMock == nil {
		s.orderStorageMock = &order_storage_mock.StorageMock{}
	}

	return s.orderStorageMock
}

func (s *suiteProvider) GetOrderStorage() order_storage.Storage {
	if s.orderStorage == nil {
		s.orderStorage = s.GetOrderStorageMock()
	}

	return s.orderStorage
}

func (s *suiteProvider) GetStockStorageMock() *stock_storage_mock.StorageMock {
	if s.stockStorageMock == nil {
		s.stockStorageMock = &stock_storage_mock.StorageMock{}
	}

	return s.stockStorageMock
}

func (s *suiteProvider) GetStockStorage() stock_storage.Storage {
	if s.stockStorage == nil {
		s.stockStorage = s.GetStockStorageMock()
	}

	return s.stockStorage
}
