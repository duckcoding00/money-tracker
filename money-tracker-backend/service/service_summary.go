package service

import (
	"context"

	"github.com/duckcoding00/money-tracker/money-tracker-backend/model/response"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/repository"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/repository/sql"
)

type SummaryService struct {
	repo *repository.Repository
}

func (s *SummaryService) GetSummary(ctx context.Context, year, month int) (*response.Summary, error) {
	userID := ctx.Value("id").(int)

	exists, err := s.repo.Sql.CheckSummary(ctx, sql.CheckSummaryParams{
		UserID: int32(userID),
		Year:   int32(year),
		Month:  int32(month),
	})

	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, ErrEmpty
	}

	result, err := s.repo.Sql.GetSummary(ctx, sql.GetSummaryParams{
		UserID: int32(userID),
		Year:   int32(year),
		Month:  int32(month),
	})

	if err != nil {
		return nil, err
	}

	return &response.Summary{
		ID:           result.ID,
		Balance:      int64(result.Balance.Int32),
		TotalIncome:  int64(result.TotalIncome),
		TotalExpense: int64(result.TotalExpense),
		Month:        int(result.Month),
		Year:         int(result.Year),
	}, nil
}
