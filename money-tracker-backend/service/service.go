package service

import (
	"context"

	"github.com/duckcoding00/money-tracker/money-tracker-backend/model/request"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/model/response"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/repository"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/utils/auth"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	User interface {
		Create(context.Context, *request.UserRequest) (int, error)
		Login(context.Context, *request.LoginRequest) (*response.LoginResponse, error)
	}

	Token interface {
		Check(token string) (*auth.JwtCustomPayload, error)
	}
}

func NewService(db *pgxpool.Pool, auth auth.JwtMethod) *Service {
	queries := repository.New(db)

	return &Service{
		User: &UserService{
			q:    queries,
			db:   db,
			auth: auth,
		},
		Token: &TokenService{
			auth: auth,
		},
	}
}
