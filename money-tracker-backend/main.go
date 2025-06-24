package main

import (
	"github.com/duckcoding00/money-tracker/money-tracker-backend/app"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	app, err := app.InitApp()
	if err != nil {
		log.Fatalf("failed to setup application : %v", err)
	}

	api := app.Config()
	if err := app.Run(api); err != nil {
		log.Fatalf("failed to start application : %v", err)
	}
}
