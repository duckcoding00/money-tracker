-- +goose Up
-- +goose StatementBegin
create table if not exists incomes(
    id text primary key,
    user_id integer not null,
    amount integer not null check (amount > 0),
    source text,
    created_at timestamptz,
    constraint fk_income_user_id foreign key (user_id) references users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists incomes;
-- +goose StatementEnd
