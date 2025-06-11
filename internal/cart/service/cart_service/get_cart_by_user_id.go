package cart_service

import (
	"cart/internal/cart/logger"
	"context"
	"sort"
)

func (s *service) GetCartByUserID(ctx context.Context, userID int64) (*Cart, error) {
	const operation = "cart_service.GetCartByUserID"

	cart, err := s.cartStorage.GetCartByUserID(ctx, userID)
	if err != nil {
		logger.Errorf(ctx, "%s: failed to get cart by user id: %s", operation, err)

		return nil, err
	}

	c := &Cart{}
	c.Items = make([]*Item, len(cart.Items))
	i := 0
	for id, quantity := range cart.Items {
		product, err := s.productService.GetProduct(ctx, id)
		if err != nil {
			logger.Errorf(ctx, "%s: failed to get product name, price: %s", operation, err)

			return nil, err
		}

		c.TotalPrice += product.Price * uint32(quantity)
		c.Items[i] = &Item{
			SkuID:    id,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: quantity,
		}

		i++
	}

	sort.SliceStable(c.Items, func(i, j int) bool {
		return c.Items[i].SkuID < c.Items[j].SkuID
	})

	return c, nil
}
