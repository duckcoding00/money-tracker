-- name: InsertExpense :one
insert into expenses (id, user_id, amount, description, category, created_at)
values ($1, $2, $3, $4, $5, $6)
returning id, user_id, amount, description, category, created_at;

-- name: UpdateExpense :one
update expenses set amount = $1, description = $2, category = $3
where id = $4
returning id, user_id, amount, description, category, created_at;

-- name: DeleteExpense :exec
delete from expenses where id = $1;

-- name: GetExpense :one
select id, user_id, amount, description, category, created_at
from expenses where id = $1;

-- name: GetExpenses :many
select id, user_id, amount, description, category, created_at
from expenses where user_id = $1
order by created_at desc;

-- name: GetExpensesByMonth :many
select id, user_id, amount, description, category, created_at
from expenses
where user_id = $1
and created_at >= date_trunc('month', $2::timestamptz)
and created_at < date_trunc('month', $2::timestamptz) + interval '1 month'
order by created_at desc;

-- name: GetExpensesByWeek :many
select id, user_id, amount, description, category, created_at
from expenses
where user_id = $1
and created_at >= date_trunc('week', $2::timestamptz)
and created_at < date_trunc('week', $2::timestamptz) + interval '7 days'
order by created_at desc;
