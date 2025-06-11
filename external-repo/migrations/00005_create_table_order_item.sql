-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.order_item
(
    order_item_id bigserial NOT NULL,
    order_id bigint NOT NULL,
    sku_id bigint NOT NULL,
    quantity integer,
    CONSTRAINT order_item_pkey PRIMARY KEY (order_item_id),
    CONSTRAINT order_item_order_id_fkey FOREIGN KEY (order_id)
        REFERENCES public."order" (order_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.order_item
    OWNER to postgres;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.order_item;
-- +goose StatementEnd
