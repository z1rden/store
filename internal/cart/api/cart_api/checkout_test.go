package cart_api_test

import (
	"bufio"
	"bytes"
	"cart/internal/cart/api/cart_api"
	"cart/internal/cart/model"
	"cart/internal/cart/suite"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiCheckout(t *testing.T) {
	type test struct {
		Name                     string
		OrderID                  int64
		RequestBody              *cart_api.CheckoutRequest
		CartServiceCheckoutError error
		StatusCode               int
	}

	var (
		CartServiceCheckoutError = errors.New("CartServiceCheckoutError")
	)

	tests := []*test{
		{
			Name:        "Неправильный формат JSON",
			RequestBody: &cart_api.CheckoutRequest{UserID: -1},
			StatusCode:  http.StatusBadRequest,
		},
		{
			Name:                     "Корзины не существует",
			RequestBody:              &cart_api.CheckoutRequest{UserID: 1},
			CartServiceCheckoutError: model.ErrNotFound,
			StatusCode:               http.StatusNotFound,
		},
		{
			Name:                     "Внутренняя ошибка",
			RequestBody:              &cart_api.CheckoutRequest{UserID: 2},
			CartServiceCheckoutError: CartServiceCheckoutError,
			StatusCode:               http.StatusInternalServerError,
		},
		{
			Name:        "Успешный тест",
			RequestBody: &cart_api.CheckoutRequest{UserID: 2},
			StatusCode:  http.StatusOK,
		},
	}

	t.Parallel()

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			sp := suite.NewSuiteProvider()
			a := cart_api.NewApi(
				sp.GetCartService())

			sp.GetCartServiceMock().EXPECT().
				Checkout(mock.Anything, tt.RequestBody.UserID).
				Return(tt.OrderID, tt.CartServiceCheckoutError)

			jsonRequest, err := json.Marshal(tt.RequestBody)
			assert.NoError(t, err, "Не должно быть ошибки во время маршалинга")

			var body bytes.Buffer
			bodyWriter := bufio.NewWriter(&body)
			_, err = bodyWriter.Write(jsonRequest)
			assert.NoError(t, err)
			err = bodyWriter.Flush()
			assert.NoError(t, err)

			reader := bufio.NewReader(&body)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/cart/checkout", reader)

			a.Checkout()(w, r)
			assert.Equal(t, tt.StatusCode, w.Result().StatusCode)
		})
	}
}
