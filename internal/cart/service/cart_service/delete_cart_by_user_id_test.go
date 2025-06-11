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

func TestServiceDeleteCartByUserId(t *testing.T) {
	type test struct {
		Name                           string
		UserID                         int64
		StorageDeleteCartByUserIDError error
		Error                          error
	}

	var (
		CartNotFoundErr = errors.New("Ошибка удаления несуществующей корзины")
	)

	var tests = []test{
		{
			Name:                           "Ошибка удаления несуществующей корзины",
			UserID:                         1,
			StorageDeleteCartByUserIDError: CartNotFoundErr,
			Error:                          CartNotFoundErr,
		},
		{
			Name:   "Успешный тест",
			UserID: 2,
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
				DeleteCartByUserID(mock.Anything, tt.UserID).
				Return(tt.StorageDeleteCartByUserIDError)

			err := s.DeleteCartByUserId(ctx, tt.UserID)
			if tt.Error != nil {
				assert.NotNil(t, tt.Error, "Должна возникать ошибка")
				assert.ErrorIs(t, err, tt.Error, "Должна соответствовать этой ошибке")
			} else {
				assert.NoError(t, err, "Ошибки не должно возникать")
			}
		})
	}
}
