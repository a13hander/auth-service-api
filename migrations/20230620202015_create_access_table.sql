-- +goose Up
-- +goose StatementBegin
CREATE TABLE access
(
    id       serial PRIMARY KEY,
    endpoint text NOT NULL,
    role     int NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS access;
-- +goose StatementEnd
