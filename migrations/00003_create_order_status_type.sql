-- +goose Up
-- +goose StatementBegin
CREATE TYPE public.order_status_type AS ENUM
    ('new', 'awaiting_payment', 'payed', 'cancelled', 'failed');

ALTER TYPE public.order_status_type
    OWNER TO postgres;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS public.order_status_type;
-- +goose StatementEnd
