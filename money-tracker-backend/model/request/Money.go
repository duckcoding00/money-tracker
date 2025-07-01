package request

type Income struct {
	Amount int32  `json:"amount" validate:"required,min=1"`
	Source string `json:"source" validate:"required"`
}

type Expense struct {
	Amount      int32  `json:"amount" validate:"required,min=1"`
	Description string `json:"description" validate:"required"`
	Category    string `json:"category" validate:"required"`
}

type UpdateIncome struct {
	Amount *int32  `json:"amount"`
	Source *string `json:"source"`
}

type UpdateExpense struct {
	Amount      *int32  `json:"amount" `
	Description *string `json:"description"`
	Category    *string `json:"category" `
}
