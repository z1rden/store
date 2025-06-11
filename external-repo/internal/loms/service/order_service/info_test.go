package order_service_test

import (
	"context"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"loms/internal/loms/model"
	"loms/internal/loms/repository/order_storage"
	"loms/internal/loms/service/order_service"
	"loms/internal/loms/suite"
	"testing"
)

func TestOrderInfo(t *testing.T) {

	type test struct {
		Name                     string
		OrderID                  int64
		Order                    *order_storage.Order
		Error                    error
		OrderStorafeGetByIDError error
	}

	tests := []*test{
		{
			Name:                     "Заказ не найден",
			OrderID:                  1,
			OrderStorafeGetByIDError: model.ErrNotFound,
			Error:                    model.ErrNotFound,
		},
		{
			Name:    "Заказ существует",
			OrderID: 2,
			Order: &order_storage.Order{
				OrderID: 2,
				UserID:  1,
				Status:  model.OrderStatusNew,
				Items: []*order_storage.Item{
					{
						SkuID:    1,
						Quantity: 1,
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
				Return(tt.Order, tt.OrderStorafeGetByIDError)

			order, err := orderService.Info(context.Background(), tt.OrderID)
			if tt.Error != nil {
				assert.NotNil(t, err, "Должна быть ошибка")
				assert.ErrorIs(t, err, tt.Error, "Не та ошибка")
			} else {
				require.Nil(t, err, "Не должно быть ошибки")
				diff := deep.Equal(order_service.ToModelOrder(tt.Order), order)
				if diff != nil {
					t.Errorf("Заказы должны совпасть: %+v", diff)
				}
			}
		})
	}
}
