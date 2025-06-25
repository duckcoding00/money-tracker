-- name: InsertSession :exec
insert into sessions (user_id, token, created_at, expired_at)
values ($1, $2, $3, $4)
on conflict (user_id)
do update set 
            token = excluded.token, 
            expired_at = excluded.expired_at, 
            created_at = excluded.created_at;

-- name: GetSessionByUserID :one
select user_id, token, created_at, expired_at
from sessions
where user_id = $1;

-- name: GetSessionByToken :one
select user_id, token, created_at, expired_at
from sessions
where token = $1;