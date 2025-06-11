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

func TestApiDeleteCartByUserID(t *testing.T) {
	type test struct {
		Name                         string
		UserID                       int64
		ServiceDeleteCartByUserIDErr error
		StatusCode                   int
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
			Name:                         "Внутренняя ошибка",
			UserID:                       1,
			StatusCode:                   http.StatusInternalServerError,
			ServiceDeleteCartByUserIDErr: errors.New("internal error"),
		},
		{
			Name:       "Успешный тест",
			UserID:     1,
			StatusCode: http.StatusNoContent,
		},
	}

	t.Parallel()

	for _, tt := range tests {
		sp := suite.NewSuiteProvider()
		a := cart_api.NewApi(
			sp.GetCartService())

		sp.GetCartServiceMock().EXPECT().
			DeleteCartByUserId(mock.Anything, tt.UserID).
			Return(tt.ServiceDeleteCartByUserIDErr)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/user/%d/cart", tt.UserID), nil)
		r.SetPathValue("user_id", fmt.Sprintf("%d", tt.UserID))

		a.DeleteCartByUserID()(w, r)

		assert.Equal(t, tt.StatusCode, w.Code)
	}
}
