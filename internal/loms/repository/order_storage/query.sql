-- name: CreateOrder :one
insert into "order"(user_id, status, created_at, updated_at)
values
    ($1,$2, now(), now())
returning order_id;

-- name: AddOrderItem :exec
insert into order_item (order_id, sku_id, quantity)
values ($1, $2, $3);

-- name: GetOrderByOrderID :one
select *
from "order"
where order_id = $1;

-- name: GetOrderItemsByOrderID :many
select *
from order_item
where
    order_id = $1;

-- name: UpdateStatusOrderByOrderID :exec
update "order"
set status = $1,
    updated_at = now()
where
    order_id = $2;