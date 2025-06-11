package order_api

import (
	"loms/internal/loms/model"
	"loms/pkg/api/order"
)

func toOrderServiceItem(item *order.OrderCreateRequest_Item) *model.Item {
	return &model.Item{
		// Есть теория, что использовать item.Sku просто дурной тон из-за особенностей ООП-языков.
		SkuID:    item.GetSku(),
		Quantity: uint16(item.GetCount()),
	}
}

func ToOrderServiceItems(items []*order.OrderCreateRequest_Item) []*model.Item {
	result := make([]*model.Item, 0, len(items))

	for _, item := range items {
		result = append(result, toOrderServiceItem(item))
	}

	return result
}

func toOrderInfoResponseItem(item *model.Item) *order.OrderInfoResponse_Item {
	return &order.OrderInfoResponse_Item{
		Sku:   item.SkuID,
		Count: uint64(item.Quantity),
	}
}

func ToOrderInfoResponseItems(items []*model.Item) []*order.OrderInfoResponse_Item {
	result := make([]*order.OrderInfoResponse_Item, 0, len(items))
	for _, item := range items {
		result = append(result, toOrderInfoResponseItem(item))
	}

	return result
}

func ToOrderInfoResponse(mOrder *model.Order) *order.OrderInfoResponse {
	return &order.OrderInfoResponse{
		Status: mOrder.Status,
		User:   mOrder.User,
		Items:  ToOrderInfoResponseItems(mOrder.Items),
	}
}
