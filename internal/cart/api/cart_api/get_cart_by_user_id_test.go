package cart_api_test

import (
	"cart/internal/cart/api/cart_api"
	"cart/internal/cart/model"
	"cart/internal/cart/service/cart_service"
	"cart/internal/cart/suite"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiGetCartByUserID(t *testing.T) {
	type test struct {
		Name                      string
		UserID                    int64
		Cart                      *cart_service.Cart
		ServiceGetCartByUserIDErr error
		StatusCode                int
	}

	tests := []test{
		{
			Name:       "UserID отрицательный",
			UserID:     -1,
			StatusCode: http.StatusBadRequest,
		},
		{
			Name:       "UserID не указан",
			UserID:     0,
			StatusCode: http.StatusBadRequest,
		},
		{
			Name:                      "Корзина для user не найдена",
			UserID:                    1,
			ServiceGetCartByUserIDErr: model.ErrNotFound,
			StatusCode:                http.StatusNotFound,
		},
		{
			Name:                      "Внутренняя ошибка",
			UserID:                    1,
			StatusCode:                http.StatusInternalServerError,
			ServiceGetCartByUserIDErr: errors.New("internal error"),
		},
		{
			Name:   "Успешный тест",
			UserID: 1,
			Cart: &cart_service.Cart{
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
				},
			},
			StatusCode: http.StatusOK,
		},
	}

	t.Parallel()

	for _, tt := range tests {
		sp := suite.NewSuiteProvider()
		a := cart_api.NewApi(
			sp.GetCartService())

		sp.GetCartServiceMock().EXPECT().
			GetCartByUserID(mock.Anything, tt.UserID).
			Return(tt.Cart, tt.ServiceGetCartByUserIDErr)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/user/%d/cart", tt.UserID), nil)
		r.SetPathValue("user_id", fmt.Sprintf("%d", tt.UserID))

		a.GetCartByUserID()(w, r)

		assert.Equal(t, tt.StatusCode, w.Code)
	}
}
