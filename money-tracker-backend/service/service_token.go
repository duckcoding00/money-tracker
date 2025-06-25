package service

import (
	"errors"

	"github.com/duckcoding00/money-tracker/money-tracker-backend/utils/auth"
)

type TokenService struct {
	auth auth.JwtMethod
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
