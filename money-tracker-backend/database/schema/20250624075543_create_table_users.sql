-- +goose Up
-- +goose StatementBegin
create table if not exists users(
    id serial primary key,
    username varchar(255) not null unique,
    email varchar(255) not null unique,
    password varchar(255) not null,
    is_active boolean default false,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
