package product_service

import (
	"cart/internal/cart/model"
	"context"
)

type Client interface {
	GetProduct(ctx context.Context, skuID int64) (*Product, error)
}

type client struct {
	storage map[int64]*Product
}

func NewClient() Client {
	return &client{
		storage: map[int64]*Product{
			1: {
				"example 1",
				10,
			},
			2: {
				"example 2",
				20,
			},
			3: {
				"example 3",
				30,
			},
			4: {
				"example 4",
				40,
			},
			5: {
				"example 5",
				50,
			},
		},
	}
}

func (c *client) GetProduct(ctx context.Context, skuID int64) (*Product, error) {
	product, exists := c.storage[skuID]
	if !exists {
		return nil, model.ErrNotFound
	}

	return product, nil
}
