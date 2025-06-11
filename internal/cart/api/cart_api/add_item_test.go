package cart_api_test

import (
	"bufio"
	"bytes"
	"cart/internal/cart/api/cart_api"
	"cart/internal/cart/model"
	"cart/internal/cart/suite"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiAddItem(t *testing.T) {
	type requestBody struct {
		Count int64
	}
	type test struct {
		Name                string
		UserID              int64
		SkuID               int64
		ServiceAddItemError error
		StatusCode          int
		RequestBody         *requestBody
	}

	tests := []test{
		{
			Name:        "UserID отрицательный",
			UserID:      -1,
			SkuID:       1,
			RequestBody: &requestBody{Count: 1},
			StatusCode:  http.StatusBadRequest,
		},
		{
			Name:        "UserID не указан",
			UserID:      0,
			SkuID:       1,
			RequestBody: &requestBody{Count: 1},
			StatusCode:  http.StatusBadRequest,
		},
		{
			Name:        "SkuID отрицательный",
			UserID:      1,
			SkuID:       -1,
			RequestBody: &requestBody{Count: 1},
			StatusCode:  http.StatusBadRequest,
		},
		{
			Name:        "SkuID не указан",
			UserID:      1,
			SkuID:       0,
			RequestBody: &requestBody{Count: 1},
			StatusCode:  http.StatusBadRequest,
		},
		{
			Name:        "Quantity отрицательный/Передан неправильный JSON",
			UserID:      1,
			SkuID:       1,
			RequestBody: &requestBody{Count: -1},
			StatusCode:  http.StatusBadRequest,
		},
		{
			Name:        "Quantity не указан",
			UserID:      1,
			SkuID:       1,
			RequestBody: &requestBody{Count: 0},
			StatusCode:  http.StatusBadRequest,
		},
		{
			Name:                "Передан несуществующий продукт",
			UserID:              1,
			SkuID:               1,
			RequestBody:         &requestBody{Count: 1},
			StatusCode:          http.StatusNotFound,
			ServiceAddItemError: model.ErrNotFound,
		},
		{
			Name:                "Внутренняя ошибка",
			UserID:              1,
			SkuID:               3,
			RequestBody:         &requestBody{Count: 1},
			StatusCode:          http.StatusInternalServerError,
			ServiceAddItemError: errors.New("internal error"),
		},
		{
			Name:        "Удачный тест",
			UserID:      1,
			SkuID:       2,
			RequestBody: &requestBody{Count: 1},
			StatusCode:  http.StatusOK,
		},
	}

	t.Parallel()

	for _, tt := range tests {
		sp := suite.NewSuiteProvider()
		a := cart_api.NewApi(sp.GetCartService())

		sp.GetCartServiceMock().EXPECT().
			AddItem(mock.Anything, tt.UserID, tt.SkuID, uint16(tt.RequestBody.Count)).
			Return(tt.ServiceAddItemError)

		jsonRequest, err := json.Marshal(tt.RequestBody)
		assert.NoError(t, err, "Не должно возникать ошибки при маршалинге")

		var body bytes.Buffer
		bodyWriter := bufio.NewWriter(&body)
		_, err = bodyWriter.Write(jsonRequest)
		assert.NoError(t, err)
		err = bodyWriter.Flush()
		assert.NoError(t, err)

		reader := bufio.NewReader(&body)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/user/%d/cart/%d", tt.UserID, tt.SkuID), reader)
		r.SetPathValue("user_id", fmt.Sprintf("%d", tt.UserID))
		r.SetPathValue("sku_id", fmt.Sprintf("%d", tt.SkuID))
		a.AddItem()(w, r)

		assert.Equal(t, tt.StatusCode, w.Result().StatusCode)
	}

}
