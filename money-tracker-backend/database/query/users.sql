-- name: InsertUser :one
insert into users (username, email, password, is_active)
values ($1, $2, $3, false)
returning id;

-- name: GetUserByID :one
select id, username, email, password, is_active, created_at, updated_at
from users 
where id = $1;

-- name: GetUserByUsername :one
select id, username, email, password, is_active, created_at, updated_at
from users 
where username = $1;

-- name: GetUserByEmail :one
select id, username, email, password, is_active, created_at, updated_at
from users 
where email = $1;

-- name: UpdateIsActive :one
update users 
set is_active = $1 
where id = $2
returning is_active;

-- name: DeleteUserByID :exec
delete from users
where id = $1;

-- name: UpdateUsername :one
update users 
set username = $1
where id = $2
returning username;

-- name: UpdatePassword :one
update users
set password = $1
where id = $2
returning password;

-- name: UpdateEmail :one
update users
set email = $1
where id = $2
returning email;