package model

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

const (
	OrderStatusNew             = "new"
	OrderStatusAwaitingPayment = "awaiting_payment"
	OrderStatusPayed           = "payed"
	OrderStatusCanceled        = "cancelled"
	OrderStatusFailed          = "failed"
)

const (
	EventOrderStatusChanged = "order_status_changed"
)

const (
	EntityTypeOrder = "order"
)

type Item struct {
	SkuID    int64
	Quantity uint16
}

type Order struct {
	OrderID int64
	User    int64
	Status  string
	Items   []*Item
}

type Orders map[int64]*Order
type OrderChangeStatusMessageOrder struct {
	OrderID int64  `json:"id"`
	Status  string `json:"status"`
}

type OrderChangeStatusMessage struct {
	ID         pgtype.UUID                  `json:"id"`
	Time       time.Time                    `json:"time"`
	Event      string                       `json:"event"`
	EntityType string                       `json:"entity_type"`
	EntityID   string                       `json:"entity_id"`
	Data       OrderChangeStatusMessageData `json:"data"`
}

type OrderChangeStatusMessageData struct {
	Order OrderChangeStatusMessageOrder `json:"order"`
}
