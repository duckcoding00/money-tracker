package app

import (
	"github.com/duckcoding00/money-tracker/money-tracker-backend/utils"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

// more setup in this file

func LoadConfig() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, err
	}

	config := Config{
		portAddress: utils.GetEnvString("PORT_ADDRESS", ""),
		dbConfig:    DBConfig{},
	}

	return config, nil
}

func InitApp() (*Application, error) {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("failed to load env", err)
	}

	return &Application{
		config: config,
	}, nil
}
