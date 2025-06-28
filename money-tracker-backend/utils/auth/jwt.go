package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var TokenTime = map[string]time.Duration{
	"access_token":  time.Minute * 10,
	"refresh_token": time.Minute * 60 * 24 * 7,
}

type JwtMethod interface {
	GeneratedToken(id int, username string, time string) (string, error)
	ValidateAccessToken(token string) (*jwt.Token, error)
	ValidateRefreshToken(token string) (*jwt.Token, error)
}

type JwtStruct struct {
	secret string
}

type JwtCustomPayload struct {
	Id       int
	Username string
	jwt.RegisteredClaims
}

func NewJwt(secret string) *JwtStruct {
	return &JwtStruct{
		secret: secret,
	}
}

func (j *JwtStruct) GeneratedToken(id int, username string, timeExp string) (string, error) {
	claims := JwtCustomPayload{
		Id:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenTime[timeExp]).UTC()),
			Subject:   strconv.Itoa(id),
			Audience:  jwt.ClaimStrings{"money-tracker-app"},
			Issuer:    "money-tracker-server",
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			NotBefore: jwt.NewNumericDate(time.Now().UTC()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", fmt.Errorf("failed generated token : %w", err)
	}

	return tokenString, nil
}

func (j *JwtStruct) ValidateAccessToken(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &JwtCustomPayload{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", t.Header["alg"])
		}
		return []byte(j.secret), nil
	})
}

func (j *JwtStruct) ValidateRefreshToken(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &JwtCustomPayload{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", t.Header["alg"])
		}
		return []byte(j.secret), nil
	},

		jwt.WithoutClaimsValidation(),
	)
}
