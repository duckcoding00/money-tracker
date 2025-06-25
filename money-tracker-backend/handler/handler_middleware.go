package handler

import (
	"context"
	"fmt"
	"strings"

	"github.com/duckcoding00/money-tracker/money-tracker-backend/service"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/utils"
	"github.com/gofiber/fiber/v2"
)

type Middleware struct {
	s *service.Service
}

func (h *Middleware) AuthMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get("Authorization")
		if token == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(utils.ApiResponse{
				Success:   false,
				Message:   "UNAUTHORIZED",
				ErrorCode: fmt.Sprintf("UNAUTHORIZED | %d", fiber.StatusUnauthorized),
				Details:   "Missing Token",
			})
		}

		parts := strings.Split(token, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(utils.ApiResponse{
				Success:   false,
				Message:   "UNAUTHORIZED",
				ErrorCode: fmt.Sprintf("UNAUTHORIZED | %d", fiber.StatusUnauthorized),
				Details:   "Malformed Token",
			})
		}

		claims, err := h.s.Token.Check(parts[1])
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(utils.ApiResponse{
				Success:   false,
				Message:   "UNAUTHORIZED",
				ErrorCode: fmt.Sprintf("UNAUTHORIZED | %d", fiber.StatusUnauthorized),
				Details:   err.Error(),
			})
		}

		c := context.WithValue(ctx.Context(), "id", claims.Id)
		c = context.WithValue(c, "username", claims.Username)

		ctx.SetUserContext(c)

		return ctx.Next()
	}
}
