package handler

import (
	"github.com/duckcoding00/money-tracker/money-tracker-backend/utils"
	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct {
}

func (h *HealthHandler) Check(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "test",
	})
}

func (h *HealthHandler) CheckUser(ctx *fiber.Ctx) error {
	c := ctx.UserContext()

	id := c.Value("id").(int32)
	username := c.Value("username").(string)
	return ctx.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Success: true,
		Message: "OMKE GAS",
		Data: fiber.Map{
			"id":       id,
			"username": username,
		},
	})
}
