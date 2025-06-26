package main

import (
	"context"

	"github.com/duckcoding00/money-tracker/money-tracker-backend/config"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/database/connection"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/handler"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/repository"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/repository/sql"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/service"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/utils"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/utils/auth"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

// more setup in this file
func LoadConfig() (config.Config, error) {
	if err := godotenv.Load(); err != nil {
		return config.Config{}, err
	}

	config := config.Config{
		PortAddress:  utils.GetEnvString("PORT_ADDRESS", "5000"),
		RedisAddress: utils.GetEnvString("REDIS_ADDRESS", ""),
		DbConfig: config.DBConfig{
			DbAddr:      utils.GetEnvString("DATABASE_URL", ""),
			MaxOpenCons: utils.GetEnvInt("DB_MAX_CONNS", 5),
			MaxIdleCons: utils.GetEnvInt("DB_MAX_IDLE", 5),
			MaxIdleTime: utils.GetEnvString("DB_MAX_TIME_IDLE", "10m"),
		},
		JwtConfig: config.JwtConfig{
			Secret: utils.GetEnvString("JWT_SECRET", ""),
		},
	}

	return config, nil
}

func InitApp() (*Application, error) {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("failed to load env : %v", err)
	}

	db, err := connection.ConnDatabase(config.DbConfig)
	if err != nil {
		log.Fatalf("failed to connect database : %v", err)
	}

	// setup repository
	sqlQueries := sql.New(db)
	repo := repository.NewRepository(sqlQueries, config.RedisAddress)

	if err := repo.Redis.Ping(context.Background()); err != nil {
		log.Fatalf("failed to connect redis : %v", err)
	}

	// setup auth
	jwt := auth.NewJwt(config.JwtConfig.Secret)

	// setup service
	service := service.NewService(db, repo, jwt)

	// setup handler
	handler := handler.NewHandler(service)

	return &Application{
		config:  config,
		handler: *handler,
	}, nil
}
