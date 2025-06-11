-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.stock
(
    sku_id bigint NOT NULL,
    total_count integer NOT NULL,
    reserved integer NOT NULL,
    CONSTRAINT stock_pkey PRIMARY KEY (sku_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.stock
    OWNER to postgres;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.stock;
-- +goose StatementEnd
