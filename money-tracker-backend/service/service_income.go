package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/duckcoding00/money-tracker/money-tracker-backend/model/request"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/model/response"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/repository"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/repository/sql"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IncomeService struct {
	repo *repository.Repository
	db   *pgxpool.Pool
}

var (
	BaseErrInsert = errors.New("insert")
	BaseErrGet    = errors.New("get")
	BaseErrUpdate = errors.New("update")

	ErrInsert = func(table string, err error) error {
		return fmt.Errorf("%w: failed to insert %s :%s", BaseErrInsert, table, err)
	}
	ErrGet = func(table string, err error) error {
		return fmt.Errorf("%w: failed to get %s :%s", BaseErrGet, table, err)
	}
	ErrUpdate = func(table string, err error) error {
		return fmt.Errorf("%w: failed to update %s :%s", BaseErrUpdate, table, err)
	}

	ErrEmpty   = errors.New("Empty Data")
	ErrInvalid = errors.New("Invalid Input")
)

func (s *IncomeService) Insert(ctx context.Context, req *request.Income) (*response.IncomeInsert, error) {
	userID := ctx.Value("id").(int)
	now := time.Now().UTC()

	year := now.Year()
	month := now.Month()

	sumID := utils.GenID("sum", userID)
	incID := utils.GenID("inc", userID)

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, ErrTx
	}
	defer tx.Rollback(ctx)

	qtx := s.repo.Sql.WithTx(tx)

	income, err := qtx.InsertIncome(ctx, sql.InsertIncomeParams{
		ID:     incID,
		UserID: int32(userID),
		Amount: req.Amount,
		Source: pgtype.Text{
			String: req.Source,
			Valid:  true,
		},
		CreatedAt: timeNow,
	})
	if err != nil {
		return nil, ErrInsert("income", err)
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

	summary, err := qtx.UpdateTotalIncome(ctx, sql.UpdateTotalIncomeParams{
		UserID: int32(userID),
		Year:   int32(year),
		Month:  int32(month),
	})
	if err != nil {
		return nil, ErrUpdate("total_income", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &response.IncomeInsert{
		Income: response.Income{
			ID:        income.ID,
			Amount:    income.Amount,
			Source:    income.Source.String,
			CreatedAt: income.CreatedAt.Time,
			Month:     int(month),
			Year:      year,
		},

		Balance: response.Balance{
			ID:          summary.ID,
			Balance:     int64(summary.Balance.Int32),
			TotalIncome: int64(summary.TotalIncome),
		},
	}, nil
}

func (s *IncomeService) Update(ctx context.Context, id string, req *request.UpdateIncome) (*response.IncomeInsert, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, ErrTx
	}
	defer tx.Rollback(ctx)

	qtx := s.repo.Sql.WithTx(tx)

	income, err := qtx.GetIncome(ctx, id)
	if err != nil {
		return nil, ErrGet("income", err)
	}

	if income.ID == "" {
		return nil, errors.New("income didnt exists")
	}

	timeIncome := income.CreatedAt.Time.UTC()
	year := timeIncome.Year()
	month := timeIncome.Month()

	if req.Amount == nil {
		req.Amount = &income.Amount
	}

	if req.Source == nil {
		req.Source = &income.Source.String
	}

	updateIncome, err := qtx.UpdateIncome(ctx, sql.UpdateIncomeParams{
		Amount: *req.Amount,
		Source: pgtype.Text{
			String: *req.Source,
			Valid:  true,
		},
		ID: id,
	})

	if err != nil {
		return nil, ErrUpdate("income", err)
	}

	exists, err := qtx.CheckSummary(ctx, sql.CheckSummaryParams{
		UserID: int32(income.UserID),
		Year:   int32(year),
		Month:  int32(month),
	})
	if err != nil {
		return nil, ErrGet("monthly_summary", err)
	}

	if !exists {
		return nil, errors.New("monthly_summary didnt exists")
	}

	summary, err := qtx.UpdateTotalIncome(ctx, sql.UpdateTotalIncomeParams{
		UserID: income.UserID,
		Year:   int32(year),
		Month:  int32(month),
	})

	if err != nil {
		return nil, ErrUpdate("total_income", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &response.IncomeInsert{
		Income: response.Income{
			ID:        income.ID,
			Amount:    updateIncome.Amount,
			Source:    updateIncome.Source.String,
			CreatedAt: income.CreatedAt.Time,
			Month:     int(month),
			Year:      year,
		},
		Balance: response.Balance{
			ID:          summary.ID,
			Balance:     int64(summary.Balance.Int32),
			TotalIncome: int64(summary.TotalIncome),
		},
	}, nil
}

func (s *IncomeService) GetIncomes(ctx context.Context, year, month int) ([]*response.Income, error) {
	userID := ctx.Value("id").(int)
	time := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

	incomes, err := s.repo.Sql.GetIncomes(ctx, int32(userID))
	if err != nil {
		return nil, err
	}

	if year != 0 {

		if month == 0 {
			return nil, ErrInvalid
		}

		incomes, err = s.repo.Sql.GetIncomesByMonth(ctx, sql.GetIncomesByMonthParams{
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

	if incomes == nil {
		return []*response.Income{}, ErrEmpty
	}

	result := []*response.Income{}

	for _, v := range incomes {
		time := v.CreatedAt.Time.UTC()
		data := response.Income{
			ID:        v.ID,
			Amount:    v.Amount,
			Source:    v.Source.String,
			CreatedAt: v.CreatedAt.Time,
			Month:     int(time.Month()),
			Year:      time.Year(),
		}
		result = append(result, &data)
	}

	return result, nil
}
