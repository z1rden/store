package order_storage

import "loms/internal/loms/repository/order_storage/sqlc"

func toOrderItems(ItemDB []sqlc.OrderItem) []*Item {
	orderItems := make([]*Item, 0, len(ItemDB))
	for _, itemDB := range ItemDB {
		orderItems = append(orderItems, &Item{
			ID:       itemDB.OrderItemID,
			SkuID:    itemDB.SkuID,
			Quantity: uint16(itemDB.Quantity.Int32),
		})
	}

	return orderItems
}
