-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN specialisation jsonb;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN specialisation;
-- +goose StatementEnd
