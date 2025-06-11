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

func TestOrderPay(t *testing.T) {

	type test struct {
		Name                           string
		OrderID                        int64
		Order                          *order_storage.Order
		Status                         string
		Error                          error
		OrderStorageGetByIDError       error
		StockStorageReserveRemoveError error
		OrderStorageSetStatusError     error
	}

	var (
		StockStorageReserveRemoveError = errors.New("Не удалось удалить резервацию продуктов")
		OrderStorageSetStatusError     = errors.New("Не удалось изменить статус заказа")
	)

	tests := []*test{
		{
			Name:                     "Заказ не найден",
			OrderID:                  1,
			OrderStorageGetByIDError: model.ErrNotFound,
			Error:                    model.ErrNotFound,
		},
		{
			Name:    "Ошибка при снятии резерва",
			OrderID: 2,
			Status:  model.OrderStatusPayed,
			Order: &order_storage.Order{
				OrderID: 2,
				UserID:  2,
				Status:  model.OrderStatusNew,
				Items: []*order_storage.Item{
					{
						SkuID:    2,
						Quantity: 2,
					},
				},
			},
			StockStorageReserveRemoveError: StockStorageReserveRemoveError,
			Error:                          StockStorageReserveRemoveError,
		},
		{
			Name:    "Ошибка при изменении статуса",
			OrderID: 3,
			Status:  model.OrderStatusPayed,
			Order: &order_storage.Order{
				OrderID: 3,
				UserID:  3,
				Status:  model.OrderStatusNew,
				Items: []*order_storage.Item{
					{
						SkuID:    3,
						Quantity: 3,
					},
				},
			},
			OrderStorageSetStatusError: OrderStorageSetStatusError,
			Error:                      OrderStorageSetStatusError,
		},
		{
			Name:    "Оплата без ошибок",
			OrderID: 4,
			Status:  model.OrderStatusPayed,
			Order: &order_storage.Order{
				OrderID: 4,
				UserID:  4,
				Status:  model.OrderStatusNew,
				Items: []*order_storage.Item{
					{
						SkuID:    4,
						Quantity: 4,
					},
				},
			},
		},
	}

	t.Parallel()

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			ctx := context.Background()
			sp := suite.NewSuiteProvider()

			orderService := order_service.NewService(
				ctx,
				sp.GetOrderStorage(),
				sp.GetStockStorage(),
			)

			sp.GetOrderStorageMock().EXPECT().
				GetByID(mock.Anything, tt.OrderID).
				Return(tt.Order, tt.OrderStorageGetByIDError)

			if tt.Order != nil {
				modelOrder := order_service.ToModelOrder(tt.Order)
				sp.GetStockStorageMock().EXPECT().
					ReserveRemove(mock.Anything, order_service.ToStockStorageItems(modelOrder.Items)).
					Return(tt.StockStorageReserveRemoveError)
			}

			sp.GetOrderStorageMock().EXPECT().
				SetStatus(mock.Anything, tt.OrderID, tt.Status).
				Return(tt.OrderStorageSetStatusError)

			err := orderService.Pay(context.Background(), tt.OrderID)
			if tt.Error != nil {
				assert.NotNil(t, err, "Должна быть ошибка")
				assert.ErrorIs(t, err, tt.Error, "Не та ошибка")
			} else {
				assert.Nil(t, err, "Не должно быть ошибки")
			}
		})
	}
}
