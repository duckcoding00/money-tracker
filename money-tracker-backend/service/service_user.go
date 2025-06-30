package service

import (
	"context"
	"errors"
	"time"

	"github.com/duckcoding00/money-tracker/money-tracker-backend/model/request"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/model/response"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/repository"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/repository/sql"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/utils/auth"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/utils/errorhandler"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.Repository
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
	db := s.repo.Sql
	// redis := s.repo.Redis
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	id, err := db.InsertUser(ctx, sql.InsertUserParams{
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
	db := s.repo.Sql

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, ErrTx
	}
	defer tx.Rollback(ctx)

	qtx := db.WithTx(tx)

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

	if err := qtx.InsertSession(ctx, sql.InsertSessionParams{
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

func (s *UserService) VerifyUser(ctx context.Context) error {
	token := ctx.Value("token").(string)

	value, err := s.repo.Redis.GetValue(ctx, token)
	if err != nil {
		return err
	}

	ttl, err := s.repo.Redis.CheckTTL(ctx, token)
	if err != nil {
		return err
	}

	if ttl <= 0 {
		return ExpiredToken
	}

	_, err = s.repo.Sql.UpdateIsActive(ctx, value)
	if err != nil {
		return err
	}

	if err := s.repo.Redis.DelValue(ctx, token); err != nil {
		return err
	}

	return nil
}

func (s *UserService) NewPassword(ctx context.Context, password string) error {
	username := ctx.Value("username").(string)
	token := ctx.Value("token").(string)

	_, err := s.repo.Redis.GetValue(ctx, token)
	if err != nil {
		return err
	}

	ttl, err := s.repo.Redis.CheckTTL(ctx, token)
	if err != nil {
		return err
	}

	if ttl <= 0 {
		return ExpiredToken
	}

	user, err := s.repo.Sql.GetUserByUsername(ctx, username)
	if err != nil {
		return err
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = s.repo.Sql.UpdatePassword(ctx, sql.UpdatePasswordParams{
		Password: string(hashPassword),
		ID:       user.ID,
	})

	if err != nil {
		return err
	}

	if err := s.repo.Redis.DelValue(ctx, token); err != nil {
		return err
	}

	return nil
}

func (s *UserService) Profile(ctx context.Context) (*sql.User, error) {
	username := ctx.Value("username").(string)

	user, err := s.repo.Sql.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
