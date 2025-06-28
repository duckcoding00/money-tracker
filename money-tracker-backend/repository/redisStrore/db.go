package redisstrore

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

var TokenDuration = map[string]time.Duration{
	"reset":         time.Minute * 5,
	"reset_session": time.Minute * 10,
	"verification":  time.Minute * 10,
}

var (
	RedisNil = errors.New("redis value didnt exists")
)

type RedisMethod interface {
	SetValue(ctx context.Context, key string, value interface{}, TokenType string) error
	GetValue(ctx context.Context, key string) (string, error)
	DelValue(ctx context.Context, key string) error
	CheckTTL(ctx context.Context, key string) (time.Duration, error)
	Ping(ctx context.Context) error
}

type Redis struct {
	client *redis.Client
}

func NewRedis(addr string) RedisMethod {
	return &Redis{
		client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: "",
			DB:       0,
		}),
	}
}

func (r *Redis) SetValue(ctx context.Context, key string, value interface{}, TokenType string) error {
	return r.client.Set(ctx, key, value, TokenDuration[TokenType]).Err()
}

func (r *Redis) GetValue(ctx context.Context, key string) (string, error) {
	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", RedisNil
		}

		return "", err
	}

	return result, nil
}

func (r *Redis) DelValue(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *Redis) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

func (r *Redis) CheckTTL(ctx context.Context, key string) (time.Duration, error) {
	return r.client.TTL(ctx, key).Result()
}
