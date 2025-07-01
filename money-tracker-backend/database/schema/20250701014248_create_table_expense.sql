-- +goose Up
-- +goose StatementBegin
create type expense_category as enum ('food_and_beverage', 'transport', 'bill', 'entertainment', 'other');

create table if not exists expenses(
    id text primary key,
    user_id integer not null,
    amount integer not null check (amount > 0),
    description text,
    category expense_category not null,
    created_at timestamptz,
    constraint fk_expense_user_id foreign key (user_id) references users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists expenses;
drop type if exists expense_category;
-- +goose StatementEnd
