package connection

import (
	"context"
	"fmt"
	"time"

	"github.com/duckcoding00/money-tracker/money-tracker-backend/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnDatabase(config config.DBConfig) (*pgxpool.Pool, error) {
	ctx := context.Background()

	configDatabase, err := pgxpool.ParseConfig(config.DbAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database address : %w", err)
	}

	duration, err := time.ParseDuration(config.MaxIdleTime)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database duration : %w", err)
	}

	configDatabase.MaxConns = int32(config.MaxIdleCons)
	configDatabase.MinConns = int32(config.MaxIdleCons)
	configDatabase.MaxConnIdleTime = duration

	dbPool, err := pgxpool.NewWithConfig(ctx, configDatabase)
	if err != nil {
		return nil, fmt.Errorf("failed to config database : %w", err)
	}

	return dbPool, nil
}
