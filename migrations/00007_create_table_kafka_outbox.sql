

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.kafka_outbox
(
    message_id uuid NOT NULL DEFAULT gen_random_uuid(),
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now(),
    status message_status_type DEFAULT 'new'::message_status_type,
    error text COLLATE pg_catalog."default",
    event text COLLATE pg_catalog."default",
    entity_type text COLLATE pg_catalog."default",
    entity_id text COLLATE pg_catalog."default",
    data text COLLATE pg_catalog."default",
    CONSTRAINT kafka_outbox_pkey PRIMARY KEY (message_id)
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.kafka_outbox
    OWNER to postgres;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.kafka_outbox;
-- +goose StatementEnd

