-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public."order"
(
    order_id bigserial NOT NULL,
    user_id bigint NOT NULL,
    status order_status_type NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now(),
    CONSTRAINT order_pkey PRIMARY KEY (order_id)
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."order"
    OWNER to postgres;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public."order";
-- +goose StatementEnd
