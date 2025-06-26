package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/duckcoding00/money-tracker/money-tracker-backend/model/request"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/model/response"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/repository"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/utils/auth"
	"github.com/gofiber/fiber/v2/log"
)

var (
	ExpiredRefreshToken = errors.New("refresh_token has expired, please login again")
	InvalidToken        = errors.New("OTP are invalid")
	ExpiredToken        = errors.New("OTP TOKEN has expired, please send again")
)

type TokenService struct {
	auth auth.JwtMethod
	repo *repository.Repository
}

func (s *TokenService) Check(token string) (*auth.JwtCustomPayload, error) {
	validToken, err := s.auth.ValidateAccessToken(token)
	if err != nil {
		return nil, err
	}

	claims, ok := validToken.Claims.(*auth.JwtCustomPayload)
	if !ok {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

func (s *TokenService) RefreshToken(token string) (*auth.JwtCustomPayload, error) {
	validToken, err := s.auth.ValidateRefreshToken(token)
	if err != nil {
		return nil, err
	}

	claims, ok := validToken.Claims.(*auth.JwtCustomPayload)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func (s *TokenService) GenerateAccessToken(ctx context.Context) (*response.LoginResponse, error) {
	db := s.repo.Sql
	id := ctx.Value("id").(int)
	username := ctx.Value("username").(string)
	now := time.Now().UTC()

	token, err := db.GetSessionByUserID(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	tokenExp := token.ExpiredAt.Time.UTC()
	log.Info(tokenExp, now, username)

	if now.After(tokenExp) {
		return nil, ExpiredRefreshToken
	}

	accessToken, err := s.auth.GeneratedToken(id, username, "access_token")
	if err != nil {
		return nil, ErrAccessToken
	}

	return &response.LoginResponse{
		Username:    username,
		AccessToken: accessToken,
	}, nil
}

func randomCode() int {
	return rand.Intn(900000) + 100000
}

func sessionCode() string {
	val := rand.Intn(900) + 100
	code := fmt.Sprintf("rst%d", val)
	return code
}

func (s *TokenService) ResetToken(ctx context.Context, username string) error {

	user, err := s.repo.Sql.GetUserByUsername(ctx, username)
	if err != nil {
		return err
	}

	value := strconv.Itoa(randomCode())
	key := fmt.Sprintf("reset:%s", user.Username)

	// TODO
	// send OTP by email or whatsapp
	return s.repo.Redis.SetValue(ctx, key, value, "reset")
}

func (s *TokenService) ValidationToken(ctx context.Context, username string) error {
	key := strconv.Itoa(randomCode())

	// TODO
	// send OTP by email or whatsapp

	return s.repo.Redis.SetValue(ctx, key, username, "verification")
}

func (s *TokenService) VerifyResetToken(ctx context.Context, req *request.VerifyToken) (string, error) {
	key := fmt.Sprintf("reset:%s", req.Username)

	value, err := s.repo.Redis.GetValue(ctx, key)
	if err != nil {
		return "", err
	}

	if value != req.Token {
		return "", InvalidToken
	}

	ttl, err := s.repo.Redis.CheckTTL(ctx, key)
	if err != nil {
		return "", err
	}

	if ttl <= 0 {
		return "", ExpiredToken
	}

	if err := s.repo.Redis.DelValue(ctx, key); err != nil {
		return "", err
	}

	sessionReset := sessionCode()
	if err := s.repo.Redis.SetValue(ctx, sessionReset, true, "reset_session"); err != nil {
		return "", err
	}

	return sessionReset, nil
}
