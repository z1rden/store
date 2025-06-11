package order_service_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"loms/internal/loms/model"
	"loms/internal/loms/service/order_service"
	"loms/internal/loms/suite"
	"testing"
)

func TestServiceCreate(t *testing.T) {
	type test struct {
		Name                       string
		UserID                     int64
		Items                      []*model.Item
		OrderID                    int64
		Status                     string
		OrderStorageCreateError    error
		StockStorageReserveError   error
		OrderStorageSetStatusError error
		Error                      error
	}

	var (
		OrderStorageCreateError    = errors.New("Возникла ошибка при создании заказа")
		StockStorageReserveError   = errors.New("Возникла ошибка при резервировании/не удалось зарезервировать")
		OrderStorageSetStatusError = errors.New("Возникла ошибка при изменении статуса заказа")
	)

	tests := []test{
		{
			Name:   "Ошибка при создании заказа",
			UserID: 1,
			Items: []*model.Item{
				{
					SkuID:    1,
					Quantity: 1,
				},
			},
			OrderStorageCreateError: OrderStorageCreateError,
			Error:                   OrderStorageCreateError,
		},
		{
			Name:   "Не удалось зарезервировать",
			UserID: 1,
			Items: []*model.Item{
				{
					SkuID:    1,
					Quantity: 1,
				},
			},
			Status:                   model.OrderStatusFailed,
			StockStorageReserveError: StockStorageReserveError,
			Error:                    StockStorageReserveError,
		},
		{
			Name:   "Не удалось изменить статус заказа на failed",
			UserID: 1,
			Items: []*model.Item{
				{
					SkuID:    1,
					Quantity: 1,
				},
			},
			Status:                     model.OrderStatusFailed,
			StockStorageReserveError:   StockStorageReserveError,
			OrderStorageSetStatusError: OrderStorageSetStatusError,
			Error:                      OrderStorageSetStatusError,
		},
		{
			Name:   "Не удалось изменить статус заказа на awayting_payment",
			UserID: 1,
			Items: []*model.Item{
				{
					SkuID:    1,
					Quantity: 1,
				},
			},
			Status:                     model.OrderStatusAwaitingPayment,
			OrderStorageSetStatusError: OrderStorageSetStatusError,
			Error:                      OrderStorageSetStatusError,
		},
		{
			Name:   "Успешный тест",
			UserID: 1,
			Items: []*model.Item{
				{
					SkuID:    1,
					Quantity: 1,
				},
			},
			Status: model.OrderStatusAwaitingPayment,
		},
	}

	t.Parallel()

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			ctx := context.Background()
			sp := suite.NewSuiteProvider()

			s := order_service.NewService(ctx,
				sp.GetOrderStorage(),
				sp.GetStockStorage())

			osItems := order_service.ToOrderStorageItems(tt.Items)
			sp.GetOrderStorageMock().EXPECT().
				Create(mock.Anything, tt.UserID, osItems).
				Return(tt.OrderID, tt.OrderStorageCreateError)

			ssItems := order_service.ToStockStorageItems(tt.Items)
			sp.GetStockStorageMock().EXPECT().
				Reserve(mock.Anything, ssItems).
				Return(tt.StockStorageReserveError)

			sp.GetOrderStorageMock().EXPECT().
				SetStatus(ctx, tt.OrderID, tt.Status).
				Return(tt.OrderStorageSetStatusError)

			orderID, err := s.Create(ctx, tt.UserID, tt.Items)
			if tt.Error != nil {
				assert.Equal(t, tt.Error, err)
			} else {
				assert.NotNil(t, orderID)
				assert.Nil(t, err)
			}
		})
	}
}
