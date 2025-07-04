package main

import (
	"fmt"

	"github.com/duckcoding00/money-tracker/money-tracker-backend/config"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/handler"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Application struct {
	config  config.Config
	handler handler.Handler
}

func (app *Application) Config() *fiber.App {
	r := fiber.New(fiber.Config{
		// catch all uncaught error
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
				Success:   false,
				Message:   "UNKNOWN_ERROR",
				ErrorCode: fmt.Sprintf("INTERNAL_STATUS_ERROR | %d", fiber.StatusInternalServerError),
				Details:   err.Error(),
			})
		},
	})

	// setup all middleware in here
	//recover
	r.Use(recover.New())
	r.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	api := r.Group("/api")
	v1 := api.Group("/v1")

	// test route
	v1.Route("/test", func(router fiber.Router) {
		router.Get("/", app.handler.Health.Check)
		router.Get("/user", app.handler.Middleware.AuthMiddleware(), app.handler.Health.CheckUser)
	})

	// user route
	v1.Route("/user", func(router fiber.Router) {
		router.Get("/", app.handler.Middleware.AuthMiddleware(), app.handler.User.Profile)
		router.Post("/", app.handler.User.Register)
		router.Post("/login", app.handler.User.Login)
		router.Patch("/reset-password", app.handler.User.ResetPassword)
		router.Patch("/verify", app.handler.User.VerifyUser)
	})

	// token route
	v1.Route("/token", func(router fiber.Router) {
		router.Post("/refresh", app.handler.Middleware.RefreshTokenMiddleware(), app.handler.Token.RefreshToken)
		router.Post("/reset", app.handler.Token.ResetToken)
		router.Post("/verify", app.handler.Token.VerifyToken)
	})

	// income route
	v1.Route("/income", func(router fiber.Router) {
		router.Post("/", app.handler.Middleware.AuthMiddleware(), app.handler.Income.InsertIncome)
		router.Patch("/:id", app.handler.Middleware.AuthMiddleware(), app.handler.Income.UpdateIncome)
		router.Get("/", app.handler.Middleware.AuthMiddleware(), app.handler.Income.GetIncomes)
	})

	// expense route
	v1.Route("/expense", func(router fiber.Router) {
		router.Post("/", app.handler.Middleware.AuthMiddleware(), app.handler.Expense.InsertExpense)
		router.Patch("/:id", app.handler.Middleware.AuthMiddleware(), app.handler.Expense.UpdateExpense)
		router.Get("/", app.handler.Middleware.AuthMiddleware(), app.handler.Expense.GetExpenses)

	})

	v1.Route("/dashboard", func(router fiber.Router) {
		router.Get("/", app.handler.Middleware.AuthMiddleware(), app.handler.Dashboard.GetDashboard)
	})

	return r
}

func (app *Application) Run(r *fiber.App) error {
	log.Infof("server running at port: %v", app.config.PortAddress)
	return r.Listen(fmt.Sprintf(":%s", app.config.PortAddress))
}
