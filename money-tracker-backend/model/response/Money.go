package response

import "time"

type IncomeInsert struct {
	Income  `json:"income"`
	Balance `json:"balance"`
}

type Income struct {
	ID        string    `json:"id"`
	Amount    int32     `json:"amount"`
	Source    string    `json:"source"`
	CreatedAt time.Time `json:"created_at"`
	Month     int       `json:"month"`
	Year      int       `json:"year"`
}

type ExpenseInsert struct {
	Expense `json:"expense"`
	Balance `json:"balance"`
}

type Expense struct {
	ID          string    `json:"id"`
	Amount      int32     `json:"amount"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	Month       int       `json:"month"`
	Year        int       `json:"year"`
}

type Balance struct {
	ID           string `json:"id"`
	Balance      int64  `json:"balance"`
	TotalIncome  int64  `json:"total_income,omitempty"`
	TotalExpense int64  `json:"total_expense,omitempty"`
}

type Summary struct {
	ID           string `json:"id"`
	Balance      int64  `json:"balance"`
	TotalIncome  int64  `json:"total_income"`
	TotalExpense int64  `json:"total_expense"`
	Month        int    `json:"month"`
	Year         int    `json:"year"`
}
