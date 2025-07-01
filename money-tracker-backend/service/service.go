package service

import (
	"context"

	"github.com/duckcoding00/money-tracker/money-tracker-backend/model/request"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/model/response"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/repository"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/repository/sql"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/utils/auth"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	User interface {
		Create(context.Context, *request.UserRequest) (int, error)
		Login(context.Context, *request.LoginRequest) (*response.LoginResponse, error)
		NewPassword(ctx context.Context, password string) error
		VerifyUser(ctx context.Context) error
		Profile(ctx context.Context) (*sql.User, error)
	}

	Token interface {
		Check(token string) (*auth.JwtCustomPayload, error)
		RefreshToken(token string) (*auth.JwtCustomPayload, error)
		GenerateAccessToken(ctx context.Context) (*response.LoginResponse, error)
		ResetToken(ctx context.Context, username string) error
		ValidationToken(ctx context.Context, username string) error
		VerifyResetToken(ctx context.Context, req *request.VerifyToken) (string, error)
	}

	Income interface {
		Insert(ctx context.Context, req *request.Income) (*response.IncomeInsert, error)
		Update(ctx context.Context, id string, req *request.UpdateIncome) (*response.IncomeInsert, error)
		GetIncomes(ctx context.Context, year, month int) ([]*response.Income, error)
	}

	Expense interface {
		Insert(ctx context.Context, req *request.Expense) (*response.ExpenseInsert, error)
		Update(ctx context.Context, id string, req *request.UpdateExpense) (*response.ExpenseInsert, error)
		GetExpenses(ctx context.Context, year, month int) ([]*response.Expense, error)
	}

	Summary interface {
		GetSummary(ctx context.Context, year, month int) (*response.Summary, error)
	}
}

func NewService(db *pgxpool.Pool, repo *repository.Repository, auth auth.JwtMethod) *Service {
	return &Service{
		User: &UserService{
			repo: repo,
			db:   db,
			auth: auth,
		},
		Token: &TokenService{
			repo: repo,
			auth: auth,
		},
		Income: &IncomeService{
			repo: repo,
			db:   db,
		},
		Expense: &ExpenseService{
			repo: repo,
			db:   db,
		},
		Summary: &SummaryService{
			repo: repo,
		},
	}
}
