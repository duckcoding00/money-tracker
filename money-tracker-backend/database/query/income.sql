-- name: InsertIncome :one
insert into incomes(id, user_id, amount, source, created_at)
values ($1, $2, $3, $4, $5)
returning id, user_id, amount, source, created_at;

-- name: UpdateIncome :one
update incomes set amount = $1, source = $2
where id = $3
returning id, user_id, amount, source, created_at;

-- name: UpdateIncomeSource :one
update incomes set source = $1
where id = $2
returning id, user_id, amount, source, created_at;

-- name: DeleteIncome :exec
delete from incomes where id = $1;

-- name: GetIncome :one
select id, user_id, amount, source, created_at
from incomes
where id = $1;

-- name: GetIncomes :many
select id, user_id, amount, source, created_at
from incomes
where user_id = $1;

-- name: GetIncomesByMonth :many
select id, user_id, amount, source, created_at
from incomes
where user_id = $1
and created_at >= date_trunc('month', $2::timestamptz)
and created_at < date_trunc('month', $2::timestamptz) + interval '1 month'
order by created_at desc;

-- name: getIncomesByWeek :many
select id, user_id, amount, source, created_at
from incomes
where user_id = $1
and created_at >= date_trunc('week', $2::timestamptz)
and created_at < date_trunc('week', $2::timestamptz) + interval '7 days'
order by created_at desc;