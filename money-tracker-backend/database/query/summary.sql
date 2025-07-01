-- name: InsertSummaryMonth :one
insert into monthly_summary(id, user_id, year, month, total_income, total_expense, created_at)
values ($1, $2, extract(year from $3::timestamptz)::int, extract(month from $3::timestamptz)::int, $4, $5, $3)
returning balance, total_income, total_expense, year, month;

-- name: GetSummary :one
select id, balance, total_income, total_expense, year, month
from monthly_summary
where user_id = $1 and year = $2 and month = $3;

-- name: CheckSummary :one
SELECT EXISTS (
  SELECT 1 FROM monthly_summary
  WHERE user_id = $1 AND year = $2 AND month = $3
);

-- name: UpdateTotalIncome :one
with income_total as (
  select coalesce(sum(amount), 0) as total
  from incomes
  where user_id = $1
    and extract(year from created_at) = $2
    and extract(month from created_at) = $3
)
update monthly_summary
set total_income = income_total.total
from income_total
where monthly_summary.user_id = $1
  and monthly_summary.year = $2
  and monthly_summary.month = $3
returning total_income, balance, id;

-- name: UpdateTotalExpense :one
with expense_total as (
  select coalesce(sum(amount), 0) as total
  from expenses
  where user_id = $1
    and extract(year from created_at) = $2
    and extract(month from created_at) = $3
)
update monthly_summary
set total_expense = expense_total.total
from expense_total
where monthly_summary.user_id = $1
  and monthly_summary.year = $2
  and monthly_summary.month = $3
returning total_expense, balance, id;
