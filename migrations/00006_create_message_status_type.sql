-- +goose Up
-- +goose StatementBegin
CREATE TYPE public.message_status_type AS ENUM
    ('new', 'sent', 'failed');

ALTER TYPE public.message_status_type
    OWNER TO postgres;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS public.message_status_type;
-- +goose StatementEnd
