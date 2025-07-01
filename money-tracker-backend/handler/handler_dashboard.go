package handler

import (
	"sync"

	"github.com/duckcoding00/money-tracker/money-tracker-backend/service"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/utils"
	"github.com/gofiber/fiber/v2"
)

type DashboardHandler struct {
	s *service.Service
}

func (h *DashboardHandler) GetDashboard(ctx *fiber.Ctx) error {
	year := ctx.QueryInt("year", 0)
	month := ctx.QueryInt("month", 0)
	c := ctx.UserContext()

	var wg sync.WaitGroup
	var incomes interface{}
	var expenses interface{}
	var summary interface{}
	var incomeErr, expenseErr, summaryErr error

	wg.Add(1)
	go func() {
		defer wg.Done()
		incomes, incomeErr = h.s.Income.GetIncomes(c, year, month)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		expenses, expenseErr = h.s.Expense.GetExpenses(c, year, month)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		summary, summaryErr = h.s.Summary.GetSummary(c, year, month)
	}()

	wg.Wait()

	if incomeErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Success: false,
			Message: "FAILED_GET_INCOMES",
			Details: incomeErr.Error(),
		})
	}

	if expenseErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Success: false,
			Message: "FAILED_GET_EXPENSES",
			Details: expenseErr.Error(),
		})
	}

	if summaryErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Success: false,
			Message: "FAILED_GET_SUMMARY",
			Details: summaryErr.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Success: true,
		Message: "SUCCESS_GET_DASHBOARD",
		Data: map[string]interface{}{
			"incomes":  incomes,
			"expenses": expenses,
			"summary":  summary,
		},
	})
}
