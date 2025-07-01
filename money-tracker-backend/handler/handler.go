package handler

import (
	"github.com/duckcoding00/money-tracker/money-tracker-backend/service"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Health interface {
		Check(ctx *fiber.Ctx) error
		CheckUser(ctx *fiber.Ctx) error
	}

	User interface {
		Register(ctx *fiber.Ctx) error
		Login(ctx *fiber.Ctx) error
		ResetPassword(ctx *fiber.Ctx) error
		VerifyUser(ctx *fiber.Ctx) error
		Profile(ctx *fiber.Ctx) error
	}

	Middleware interface {
		AuthMiddleware() fiber.Handler
		RefreshTokenMiddleware() fiber.Handler
	}

	Token interface {
		RefreshToken(ctx *fiber.Ctx) error
		ResetToken(ctx *fiber.Ctx) error
		VerifyToken(ctx *fiber.Ctx) error
	}

	Income interface {
		InsertIncome(ctx *fiber.Ctx) error
		UpdateIncome(ctx *fiber.Ctx) error
		GetIncomes(ctx *fiber.Ctx) error
	}
	Expense interface {
		InsertExpense(ctx *fiber.Ctx) error
		UpdateExpense(ctx *fiber.Ctx) error
		GetExpenses(ctx *fiber.Ctx) error
	}
	Dashboard interface {
		GetDashboard(ctx *fiber.Ctx) error
	}
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		Health: &HealthHandler{},
		User: &UserHandler{
			s: service,
		},
		Middleware: &Middleware{
			s: service,
		},
		Token: &TokenHandler{
			s: service,
		},
		Income: &IncomeHandler{
			s: service,
		},
		Expense: &ExpenseHandler{
			s: service,
		},
		Dashboard: &DashboardHandler{
			s: service,
		},
	}
}
