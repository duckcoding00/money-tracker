package service

import (
	"context"
	"errors"
	"time"

	"github.com/duckcoding00/money-tracker/money-tracker-backend/model/request"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/model/response"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/repository"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/utils/auth"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/utils/errorhandler"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	q    *repository.Queries
	db   *pgxpool.Pool
	auth auth.JwtMethod
}

var (
	timeNow = pgtype.Timestamptz{
		Time:             time.Now().UTC(),
		Valid:            true,
		InfinityModifier: pgtype.Finite,
	}
)

var (
	ErrCredentials  = errors.New("invalid email or password")
	ErrTx           = errors.New("failed to start transactions")
	ErrAccessToken  = errors.New("failed to create access token")
	ErrRefreshToken = errors.New("failed to create refresh token")
)

func (s *UserService) Create(ctx context.Context, req *request.UserRequest) (int, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	id, err := s.q.InsertUser(ctx, repository.InsertUserParams{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashPassword),
	})

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return 0, errorhandler.HandleDuplicateError(err, req.Email, req.Username)
		}
		return 0, err
	}

	return int(id), nil
}

func (s *UserService) Login(ctx context.Context, req *request.LoginRequest) (*response.LoginResponse, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, ErrTx
	}
	defer tx.Rollback(ctx)

	qtx := s.q.WithTx(tx)

	result, err := qtx.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, ErrCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(req.Password)); err != nil {
		return nil, ErrCredentials
	}

	accessToken, err := s.auth.GeneratedToken(int(result.ID), result.Username, "access_token")
	if err != nil {
		return nil, ErrAccessToken
	}

	refreshToken, err := s.auth.GeneratedToken(int(result.ID), result.Username, "refresh_token")
	if err != nil {
		return nil, ErrRefreshToken
	}

	if err := qtx.InsertSession(ctx, repository.InsertSessionParams{
		UserID:    result.ID,
		Token:     refreshToken,
		CreatedAt: timeNow,
		ExpiredAt: pgtype.Timestamptz{
			Time:             time.Now().Add(auth.TokenTime["refresh_token"]).UTC(),
			Valid:            true,
			InfinityModifier: pgtype.Finite,
		},
	}); err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &response.LoginResponse{
		Username:    result.Username,
		AccessToken: accessToken,
	}, nil
}
