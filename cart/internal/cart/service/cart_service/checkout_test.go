package cart_service_test

import (
	"cart/internal/cart/client/loms_service"
	"cart/internal/cart/repository/cart_storage"
	"cart/internal/cart/service/cart_service"
	"cart/internal/cart/suite"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestServiceCheckout(t *testing.T) {
	type test struct {
		Name                               string
		UserID                             int64
		CartStorageGetCartByUserIDError    error
		LomsServiceOrderCreateError        error
		CartStorageDeleteCartByUserIDError error
		Cart                               *cart_storage.Cart
		OrderID                            int64
	}

	var (
		CartStorageGetCartByUserIDError    = errors.New("failed to get cart by user ID")
		LomsServiceOrderCreateError        = errors.New("failed to create order")
		CartStorageDeleteCartByUserIDError = errors.New("failed to delete order")
	)

	tests := []test{
		{
			Name:                            "Корзины у пользователя не существует",
			UserID:                          1,
			CartStorageGetCartByUserIDError: CartStorageGetCartByUserIDError,
		},
		{
			Name:   "Ошибка во время создания заказа",
			UserID: 1,
			Cart: &cart_storage.Cart{
				Items: map[int64]uint16{
					1: 1,
				},
			},
			LomsServiceOrderCreateError: LomsServiceOrderCreateError,
		},
		{
			Name:   "Ошибка при удалении корзины пользователя",
			UserID: 1,
			Cart: &cart_storage.Cart{
				Items: map[int64]uint16{
					1: 1,
				},
			},
			CartStorageDeleteCartByUserIDError: CartStorageDeleteCartByUserIDError,
		},
		{
			Name:   "Успешный тест",
			UserID: 1,
			Cart: &cart_storage.Cart{
				Items: map[int64]uint16{
					1: 1,
				},
			},
			OrderID: 1,
		},
	}

	t.Parallel()

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			sp := suite.NewSuiteProvider()

			s := cart_service.NewService(
				sp.GetCartStorage(),
				sp.GetProductService(),
				sp.GetLomsService(),
			)

			sp.GetCartStorageMock().EXPECT().
				GetCartByUserID(mock.Anything, tt.UserID).
				Return(tt.Cart, tt.CartStorageGetCartByUserIDError)

			var items []*loms_service.OrderItem
			if tt.Cart != nil {
				items = cart_service.ToOrderItems(tt.Cart.Items)
			}

			sp.GetLomsServiceMock().EXPECT().
				OrderCreate(mock.Anything, tt.UserID, items).
				Return(tt.OrderID, tt.LomsServiceOrderCreateError)

			sp.GetCartStorageMock().EXPECT().
				DeleteCartByUserID(mock.Anything, tt.UserID).
				Return(tt.CartStorageDeleteCartByUserIDError)

			orderID, err := s.Checkout(context.Background(), tt.UserID)
			if tt.Cart == nil {
				assert.ErrorIs(t, err, tt.CartStorageGetCartByUserIDError, "Должна быть ошибка, "+
					"что корзины у пользователя не существует")
			} else if tt.LomsServiceOrderCreateError != nil {
				assert.ErrorIs(t, err, tt.LomsServiceOrderCreateError, "Должна быть ошибка, связанная"+
					"с созданием заказа")
			} else if tt.CartStorageDeleteCartByUserIDError != nil {
				assert.ErrorIs(t, err, tt.CartStorageDeleteCartByUserIDError, "Должна быть ошибка, "+
					"связанная с удалением корзины пользователя")
			} else {
				assert.Equal(t, tt.OrderID, orderID)
				assert.NoError(t, err, "Не должно возникать ошибки")
			}
		})
	}
}
