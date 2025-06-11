-- +goose Up
-- +goose StatementBegin
insert into stock(sku_id, total_count, reserved)
values
    (1, 150, 10),
    (2, 200, 20),
    (3, 250, 30),
    (4, 300, 40),
    (5, 350, 50);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
truncate table stock;
-- +goose StatementEnd
