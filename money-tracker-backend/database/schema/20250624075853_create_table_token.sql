-- +goose Up
-- +goose StatementBegin
create table if not exists token(
    token text unique,
    expired_at timestamptz
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists token;
-- +goose StatementEnd
