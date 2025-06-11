package stock_service_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"loms/internal/loms/model"
	"loms/internal/loms/service/stock_service"
	"loms/internal/loms/suite"
	"testing"
)

func TestStockInfo(t *testing.T) {

	type test struct {
		Name     string
		SkuID    int64
		Quantity uint16
		Error    error
	}

	tests := []*test{
		{
			Name:  "Стока нет",
			SkuID: 1,
			Error: model.ErrNotFound,
		},
		{
			Name:     "Сток есть",
			SkuID:    2,
			Quantity: 10,
		},
	}

	t.Parallel()

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			ctx := context.Background()
			sp := suite.NewSuiteProvider()

			stockService := stock_service.NewService(
				ctx,
				sp.GetStockStorage(),
			)

			sp.GetStockStorageMock().EXPECT().
				GetBySku(mock.Anything, tt.SkuID).
				Return(tt.Quantity, tt.Error)

			quantity, err := stockService.Info(context.Background(), tt.SkuID)
			if tt.Error != nil {
				assert.NotNil(t, err, "Должна быть ошибка")
				assert.ErrorIs(t, err, tt.Error, "Не та ошибка")
			} else {
				assert.Equal(t, tt.Quantity, quantity, "Не совпало количество")
			}
		})
	}

}
