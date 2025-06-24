-- +goose Up
-- +goose StatementBegin
create table if not exists sessions(
    user_id integer unique not null,
    token text not null,
    created_at timestamp default current_timestamp,
    expired_at timestamp,
    constraint fk_session_user_id foreign key (user_id) references users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists sessions;
-- +goose StatementEnd
