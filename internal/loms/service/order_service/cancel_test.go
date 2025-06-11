package order_service_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"loms/internal/loms/model"
	"loms/internal/loms/repository/order_storage"
	"loms/internal/loms/service/order_service"
	"loms/internal/loms/suite"
	"testing"
)

func TestServiceCancel(t *testing.T) {
	type test struct {
		Name                           string
		OrderID                        int64
		OrderStorage                   *order_storage.Order
		Status                         string
		OrderStorageGetByIDError       error
		StockStorageReserveCancelError error
		OrderStorageSetStatusError     error
		Error                          error
	}

	var (
		OrderStorageGetByIDError       = errors.New("Ошибка при получении информации по заказу")
		StockStorageReserveCancelError = errors.New("Ошибка при отмене зарезервированных товаров")
		OrderStorageSetStatusError     = errors.New("Ошибка при изменении статуса заказа")
	)

	tests := []test{
		{
			Name:                     "Такого заказа не существует",
			OrderID:                  1,
			OrderStorageGetByIDError: OrderStorageGetByIDError,
			Error:                    OrderStorageGetByIDError,
		},
		{
			Name:    "Ошибка во время отмены зарезирвированных продуктов",
			OrderID: 1,
			OrderStorage: &order_storage.Order{
				OrderID: 1,
				UserID:  1,
				Status:  "xxx",
				Items: []*order_storage.Item{
					{
						SkuID:    1,
						Quantity: 1,
					},
				},
			},
			StockStorageReserveCancelError: StockStorageReserveCancelError,
			Error:                          StockStorageReserveCancelError,
		},
		{
			Name:    "Ошибка при изменении статуса заказа",
			OrderID: 1,
			OrderStorage: &order_storage.Order{
				OrderID: 1,
				UserID:  1,
				Status:  "xxx",
				Items: []*order_storage.Item{
					{
						SkuID:    1,
						Quantity: 1,
					},
				},
			},
			Status:                     model.OrderStatusCanceled,
			OrderStorageSetStatusError: OrderStorageSetStatusError,
			Error:                      OrderStorageSetStatusError,
		},
		{
			Name:    "Успешный тест",
			OrderID: 1,
			OrderStorage: &order_storage.Order{
				OrderID: 1,
				UserID:  1,
				Status:  "xxx",
				Items: []*order_storage.Item{
					{
						SkuID:    1,
						Quantity: 1,
					},
				},
			},
			Status: model.OrderStatusCanceled,
		},
	}

	t.Parallel()

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			ctx := context.Background()
			sp := suite.NewSuiteProvider()

			s := order_service.NewService(
				ctx,
				sp.GetOrderStorage(),
				sp.GetStockStorage())

			sp.GetOrderStorageMock().EXPECT().
				GetByID(mock.Anything, tt.OrderID).
				Return(tt.OrderStorage, tt.OrderStorageGetByIDError)

			if tt.OrderStorage != nil {
				mOrder := order_service.ToModelOrder(tt.OrderStorage)
				stockStorageItems := order_service.ToStockStorageItems(mOrder.Items)
				sp.GetStockStorageMock().EXPECT().
					ReserveCancel(mock.Anything, stockStorageItems).
					Return(tt.StockStorageReserveCancelError)
			}

			sp.GetOrderStorageMock().EXPECT().
				SetStatus(mock.Anything, tt.OrderID, tt.Status).
				Return(tt.OrderStorageSetStatusError)

			err := s.Cancel(ctx, tt.OrderID)
			if tt.Error != nil {
				assert.Equal(t, err, tt.Error, "Ошибки должны совпадать")
			} else {
				assert.Nil(t, err, "Не должно быть ошибки")
			}
		})
	}
}
