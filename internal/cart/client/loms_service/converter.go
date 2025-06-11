package loms_service

import (
	"cart/pkg/api/order"
	"cart/pkg/api/stock"
)

func ToOrderCreateRequest(user int64, items []*OrderItem) *order.OrderCreateRequest {
	return &order.OrderCreateRequest{
		User:  user,
		Items: toOrderCreateRequestItems(items),
	}
}

func toOrderCreateRequestItems(items []*OrderItem) []*order.OrderCreateRequest_Item {
	res := make([]*order.OrderCreateRequest_Item, 0, len(items))
	for _, item := range items {
		res = append(res, toOrderCreateRequestItem(item))
	}
	return res
}

func toOrderCreateRequestItem(item *OrderItem) *order.OrderCreateRequest_Item {
	return &order.OrderCreateRequest_Item{
		Sku:   item.Sku,
		Count: uint64(item.Quantity),
	}
}

func toStockInfoRequest(SkuID int64) *stock.StockInfoRequest {
	return &stock.StockInfoRequest{
		Sku: SkuID,
	}
}
