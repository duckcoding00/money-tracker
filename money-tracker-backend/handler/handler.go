package handler

import (
	"github.com/duckcoding00/money-tracker/money-tracker-backend/service"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/utils/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	Health interface {
		Check(ctx *fiber.Ctx) error
		CheckUser(ctx *fiber.Ctx) error
	}

	User interface {
		Register(ctx *fiber.Ctx) error
		Login(ctx *fiber.Ctx) error
	}

	Middleware interface {
		AuthMiddleware() fiber.Handler
	}
}

func NewHandler(db *pgxpool.Pool, auth auth.JwtMethod) *Handler {
	service := service.NewService(db, auth)
	return &Handler{
		Health: &HealthHandler{},
		User: &UserHandler{
			s: service,
		},
		Middleware: &Middleware{
			s: service,
		},
	}
}
