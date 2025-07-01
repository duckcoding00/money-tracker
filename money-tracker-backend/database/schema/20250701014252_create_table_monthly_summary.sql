-- +goose Up
-- +goose StatementBegin
create table if not exists monthly_summary(
    id text primary key,
    user_id integer not null,
    year integer not null,
    month integer not null,
    total_income integer not null default 0,
    total_expense integer not null default 0,
    balance integer generated always as (total_income - total_expense) stored,
    created_at timestamptz,
    unique(user_id, year, month)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists monthly_summary;
-- +goose StatementEnd
