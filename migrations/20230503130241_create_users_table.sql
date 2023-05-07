-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id         serial PRIMARY KEY,
    email      text      NOT NULL UNIQUE,
    username   text      NOT NULL,
    password   text      NOT NULL,
    role       int       NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
