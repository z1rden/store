package order_service

import (
	"loms/internal/loms/model"
	"loms/internal/loms/repository/order_storage"
	"loms/internal/loms/repository/stock_storage"
)

func ToOrderStorageItems(items []*model.Item) []*order_storage.Item {
	res := make([]*order_storage.Item, 0, len(items))
	for _, item := range items {
		res = append(res, toOrderStorageItem(item))
	}

	return res
}

func toOrderStorageItem(item *model.Item) *order_storage.Item {
	return &order_storage.Item{
		SkuID:    item.SkuID,
		Quantity: item.Quantity,
	}
}

func toStockStorageItem(item *model.Item) *stock_storage.ReserveItem {
	return &stock_storage.ReserveItem{
		SkuID:    item.SkuID,
		Quantity: item.Quantity,
	}
}

func ToStockStorageItems(items []*model.Item) []*stock_storage.ReserveItem {
	reserveItems := make([]*stock_storage.ReserveItem, 0, len(items))
	for _, item := range items {
		reserveItems = append(reserveItems, toStockStorageItem(item))
	}

	return reserveItems
}

func toModelItem(item *order_storage.Item) *model.Item {
	return &model.Item{
		SkuID:    item.SkuID,
		Quantity: item.Quantity,
	}
}

func toModelItems(items []*order_storage.Item) []*model.Item {
	res := make([]*model.Item, 0, len(items))
	for _, item := range items {
		res = append(res, toModelItem(item))
	}

	return res
}

func ToModelOrder(order *order_storage.Order) *model.Order {
	return &model.Order{
		OrderID: order.OrderID,
		User:    order.UserID,
		Status:  order.Status,
		Items:   toModelItems(order.Items),
	}
}
