-- name: GetBySku :one
select sku_id, total_count, reserved
from stock
where
    sku_id = $1 limit 1;

-- name: Reserve :exec
update stock
    set reserved = reserved + $1
where
    sku_id = $2
returning *;

-- name: ReserveCancel :exec
update stock
    set reserved = reserved - $1
where sku_id = $2
returning *;

-- name: ReserveRemove :exec
update stock
    set reserved = reserved - $1,
        total_count = total_count - $1
where
    sku_id = $2
returning *;