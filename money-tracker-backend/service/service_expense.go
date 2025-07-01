package service

import (
	"context"
	"errors"
	"time"

	"github.com/duckcoding00/money-tracker/money-tracker-backend/model/request"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/model/response"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/repository"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/repository/sql"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ExpenseService struct {
	repo *repository.Repository
	db   *pgxpool.Pool
}

func (s *ExpenseService) Insert(ctx context.Context, req *request.Expense) (*response.ExpenseInsert, error) {
	userID := ctx.Value("id").(int)
	now := time.Now().UTC()

	year := now.Year()
	month := now.Month()

	sumID := utils.GenID("sum", userID)
	expID := utils.GenID("exp", userID)

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, ErrTx
	}
	defer tx.Rollback(ctx)
	qtx := s.repo.Sql.WithTx(tx)

	expense, err := qtx.InsertExpense(ctx, sql.InsertExpenseParams{
		ID:     expID,
		UserID: int32(userID),
		Amount: req.Amount,
		Description: pgtype.Text{
			String: req.Description,
			Valid:  true,
		},
		Category:  sql.ExpenseCategory(req.Category),
		CreatedAt: timeNow,
	})

	if err != nil {
		return nil, ErrInsert("expense", err)
	}

	exists, err := qtx.CheckSummary(ctx, sql.CheckSummaryParams{
		UserID: int32(userID),
		Year:   int32(year),
		Month:  int32(month),
	})
	if err != nil {
		return nil, ErrGet("monthly_summary", err)
	}

	if !exists {
		_, err = qtx.InsertSummaryMonth(ctx, sql.InsertSummaryMonthParams{
			ID:           sumID,
			UserID:       int32(userID),
			CreatedAt:    timeNow,
			TotalIncome:  0,
			TotalExpense: 0,
		})
		if err != nil {
			return nil, ErrInsert("monthly_summary", err)
		}
	}

	summary, err := qtx.UpdateTotalExpense(ctx, sql.UpdateTotalExpenseParams{
		UserID: int32(userID),
		Year:   int32(year),
		Month:  int32(month),
	})
	if err != nil {
		return nil, ErrUpdate("total_expense", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &response.ExpenseInsert{
		Expense: response.Expense{
			ID:          expense.ID,
			Amount:      expense.Amount,
			Description: expense.Description.String,
			Category:    string(expense.Category),
			CreatedAt:   expense.CreatedAt.Time,
			Month:       int(month),
			Year:        year,
		},
		Balance: response.Balance{
			ID:           summary.ID,
			Balance:      int64(summary.Balance.Int32),
			TotalExpense: int64(summary.TotalExpense),
		},
	}, nil
}

func (s *ExpenseService) Update(ctx context.Context, id string, req *request.UpdateExpense) (*response.ExpenseInsert, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, ErrTx
	}
	defer tx.Rollback(ctx)

	qtx := s.repo.Sql.WithTx(tx)

	expense, err := qtx.GetExpense(ctx, id)
	if err != nil {
		return nil, ErrGet("income", err)
	}

	if expense.ID == "" {
		return nil, errors.New("income didnt exists")
	}

	timeExpense := expense.CreatedAt.Time.UTC()
	year := timeExpense.Year()
	month := timeExpense.Month()

	if req.Amount == nil {
		req.Amount = &expense.Amount
	}

	if req.Description == nil {
		req.Description = &expense.Description.String
	}

	if req.Category == nil {
		req.Category = (*string)(&expense.Category)
	}

	updateExpense, err := qtx.UpdateExpense(ctx, sql.UpdateExpenseParams{
		Amount: *req.Amount,
		Description: pgtype.Text{
			String: *req.Description,
			Valid:  true,
		},
		Category: sql.ExpenseCategory(*req.Category),
		ID:       id,
	})

	if err != nil {
		return nil, ErrUpdate("income", err)
	}

	exists, err := qtx.CheckSummary(ctx, sql.CheckSummaryParams{
		UserID: int32(expense.UserID),
		Year:   int32(year),
		Month:  int32(month),
	})
	if err != nil {
		return nil, ErrGet("monthly_summary", err)
	}

	if !exists {
		return nil, errors.New("monthly_summary didnt exists")
	}

	summary, err := qtx.UpdateTotalExpense(ctx, sql.UpdateTotalExpenseParams{
		UserID: expense.UserID,
		Year:   int32(year),
		Month:  int32(month),
	})

	if err != nil {
		return nil, ErrUpdate("total_income", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &response.ExpenseInsert{
		Expense: response.Expense{
			ID:          expense.ID,
			Amount:      updateExpense.Amount,
			Description: updateExpense.Description.String,
			Category:    string(updateExpense.Category),
			CreatedAt:   expense.CreatedAt.Time,
			Month:       int(month),
			Year:        year,
		},
		Balance: response.Balance{
			ID:           summary.ID,
			Balance:      int64(summary.Balance.Int32),
			TotalExpense: int64(summary.TotalExpense),
		},
	}, nil
}

func (s *ExpenseService) GetExpenses(ctx context.Context, year, month int) ([]*response.Expense, error) {
	userID := ctx.Value("id").(int)
	time := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

	expenses, err := s.repo.Sql.GetExpenses(ctx, int32(userID))
	if err != nil {
		return nil, err
	}

	if year != 0 {

		if month == 0 {
			return nil, ErrInvalid
		}

		expenses, err = s.repo.Sql.GetExpensesByMonth(ctx, sql.GetExpensesByMonthParams{
			UserID: int32(userID),
			Column2: pgtype.Timestamptz{
				Time:  time,
				Valid: true,
			},
		})

		if err != nil {
			return nil, err
		}
	}

	if expenses == nil {
		return []*response.Expense{}, ErrEmpty
	}

	result := []*response.Expense{}

	for _, v := range expenses {
		time := v.CreatedAt.Time.UTC()
		data := response.Expense{
			ID:          v.ID,
			Amount:      v.Amount,
			Description: v.Description.String,
			Category:    string(v.Category),
			CreatedAt:   v.CreatedAt.Time,
			Month:       int(time.Month()),
			Year:        time.Year(),
		}
		result = append(result, &data)
	}

	return result, nil
}
