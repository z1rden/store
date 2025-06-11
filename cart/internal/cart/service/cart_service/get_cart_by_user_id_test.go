package cart_service_test

import (
	"cart/internal/cart/client/product_service"
	"cart/internal/cart/model"
	"cart/internal/cart/repository/cart_storage"
	"cart/internal/cart/service/cart_service"
	"cart/internal/cart/suite"
	"context"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestServiceGetCartByUserID(t *testing.T) {
	type testProduct struct {
		SkuID   int64
		Product *product_service.Product
	}

	type test struct {
		Name                        string
		UserID                      int64
		CartStorage                 *cart_storage.Cart
		CartService                 *cart_service.Cart
		StorageGetCartByUserIDError error
		ProductsClient              []*testProduct
		ClientGetProductError       error
		Error                       error
	}

	tests := []test{
		{
			Name:                        "Корзины не существует",
			UserID:                      1,
			StorageGetCartByUserIDError: model.ErrNotFound,
			Error:                       model.ErrNotFound,
		},
		{
			Name:                        "Продукта не существует",
			UserID:                      1,
			StorageGetCartByUserIDError: model.ErrNotFound,
			Error:                       model.ErrNotFound,
		},
		{
			Name:   "Успешный тест",
			UserID: 1,
			CartStorage: &cart_storage.Cart{
				Items: map[int64]uint16{
					1: 1,
					2: 2,
					3: 3,
				},
			},
			CartService: &cart_service.Cart{
				Items: []*cart_service.Item{
					{
						SkuID:    1,
						Name:     "example 1",
						Quantity: 1,
						Price:    10,
					},
					{
						SkuID:    2,
						Name:     "example 2",
						Quantity: 2,
						Price:    20,
					},
					{
						SkuID:    3,
						Name:     "example 3",
						Quantity: 3,
						Price:    30,
					},
				},
				TotalPrice: 140,
			},
			ProductsClient: []*testProduct{
				{
					SkuID: 1,
					Product: &product_service.Product{
						Name:  "example 1",
						Price: 10,
					},
				},
				{
					SkuID: 2,
					Product: &product_service.Product{
						Name:  "example 2",
						Price: 20,
					},
				},
				{
					SkuID: 3,
					Product: &product_service.Product{
						Name:  "example 3",
						Price: 30,
					},
				},
			},
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
				GetCartByUserID(mock.Anything, tt.UserID).
				Return(tt.CartStorage, tt.StorageGetCartByUserIDError)

			for _, product := range tt.ProductsClient {
				sp.GetProductServiceMock().EXPECT().
					GetProduct(mock.Anything, product.SkuID).
					Return(product.Product, tt.ClientGetProductError)
			}

			cart, err := s.GetCartByUserID(ctx, tt.UserID)
			if tt.Error != nil {
				assert.ErrorIs(t, err, tt.Error, "Ошибки должны совпадать")
			} else {
				assert.NoError(t, err, "Ошибки не должно быть")
			}

			diff := deep.Equal(tt.CartService, cart)
			if diff != nil {
				t.Errorf("Корзины должны совпадать")
			}

		})
	}
}
