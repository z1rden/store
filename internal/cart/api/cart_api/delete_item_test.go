package cart_api_test

import (
	"cart/internal/cart/api/cart_api"
	"cart/internal/cart/suite"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiDeleteItem(t *testing.T) {
	type test struct {
		Name                   string
		UserID                 int64
		SkuID                  int64
		ServiceDeleteItemError error
		StatusCode             int
	}

	var tests = []test{
		{
			Name:       "UserID отрицательный",
			UserID:     -1,
			SkuID:      1,
			StatusCode: http.StatusBadRequest,
		},
		{
			Name:       "UserID не указан",
			UserID:     0,
			SkuID:      1,
			StatusCode: http.StatusBadRequest,
		},
		{
			Name:       "SkuID отрицательный",
			UserID:     1,
			SkuID:      -1,
			StatusCode: http.StatusBadRequest,
		},
		{
			Name:       "SkuID не указан",
			UserID:     1,
			SkuID:      0,
			StatusCode: http.StatusBadRequest,
		},
		{
			Name:                   "Внутренняя ошибка",
			UserID:                 1,
			SkuID:                  3,
			StatusCode:             http.StatusInternalServerError,
			ServiceDeleteItemError: errors.New("internal error"),
		},
		{
			Name:       "Успешный тест",
			UserID:     1,
			SkuID:      1,
			StatusCode: http.StatusNoContent,
		},
	}

	t.Parallel()

	for _, tt := range tests {
		sp := suite.NewSuiteProvider()
		a := cart_api.NewApi(sp.GetCartService())

		sp.GetCartServiceMock().EXPECT().
			DeleteItem(mock.Anything, tt.UserID, tt.SkuID).
			Return(tt.ServiceDeleteItemError)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/user/%d/cart/%d", tt.UserID, tt.SkuID), nil)
		r.SetPathValue("user_id", fmt.Sprintf("%d", tt.UserID))
		r.SetPathValue("sku_id", fmt.Sprintf("%d", tt.SkuID))
		a.DeleteItem()(w, r)

		assert.Equal(t, tt.StatusCode, w.Code)
	}
}
