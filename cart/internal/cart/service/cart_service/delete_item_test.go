package cart_service_test

import (
	"cart/internal/cart/service/cart_service"
	"cart/internal/cart/suite"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestServiceDeleteItem(t *testing.T) {
	type test struct {
		Name                   string
		UserID                 int64
		SkuID                  int64
		StorageDeleteItemError error
		Error                  error
	}

	var (
		CartNotFoundErr = errors.New("Корзины не существует")
		ItemNotFoundErr = errors.New("Товара не существует")
	)

	var tests = []test{
		{
			Name:                   "Корзины не существует",
			UserID:                 1,
			SkuID:                  1,
			StorageDeleteItemError: CartNotFoundErr,
			Error:                  CartNotFoundErr,
		},
		{
			Name:                   "Товара не существует",
			UserID:                 1,
			SkuID:                  1,
			StorageDeleteItemError: ItemNotFoundErr,
			Error:                  ItemNotFoundErr,
		},
		{
			Name:   "Успешный тест",
			UserID: 1,
			SkuID:  1,
		},
	}

	t.Parallel()

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			sp := suite.NewSuiteProvider()
			s := cart_service.NewService(
				sp.GetCartStorage(),
				sp.GetProductService(),
				sp.GetLomsServiceMock(),
			)
			ctx := context.Background()

			sp.GetCartStorageMock().EXPECT().
				DeleteItem(mock.Anything, tt.UserID, tt.SkuID).
				Return(tt.StorageDeleteItemError)

			err := s.DeleteItem(ctx, tt.UserID, tt.SkuID)
			if tt.Error != nil {
				assert.NotNil(t, tt.Error, "Должна возникать ошибка")
				assert.ErrorIs(t, err, tt.Error, "Должна соответствовать этой ошибке")
			} else {
				assert.NoError(t, err, "Ошибки не должно возникать.")
			}
		})
	}
}
