-- +goose Up
-- +goose StatementBegin
create index idx_income_user_date on incomes(user_id, created_at);
create index idx_expense_user_date on expenses(user_id, created_at);
create index idx_summary_monthly_month on monthly_summary(user_id, year, month);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index if exists idx_summary_monthly_month;
drop index if exists idx_expense_user_date;
drop index if exists idx_income_user_date;
-- +goose StatementEnd
