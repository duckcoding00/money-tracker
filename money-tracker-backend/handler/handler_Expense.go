package handler

import (
	"errors"
	"fmt"

	"github.com/duckcoding00/money-tracker/money-tracker-backend/model/request"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/service"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/utils"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/utils/errorhandler"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type ExpenseHandler struct {
	s *service.Service
}

func (h *ExpenseHandler) InsertExpense(ctx *fiber.Ctx) error {
	c := ctx.UserContext()

	request := new(request.Expense)

	if err := ctx.BodyParser(request); err != nil {
		log.Errorf("failed to parse request : %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
			Success:   false,
			Message:   "FAILED_PARSE_REQUES",
			ErrorCode: fmt.Sprintf("BAD_REQUEST | %d", fiber.StatusBadRequest),
			Details:   err,
		})
	}

	if err := request.Validate(); err != nil {
		validationErrs := errorhandler.ValidationErrors(err.(validator.ValidationErrors))
		log.Errorf("validation errors : %v", validationErrs)
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
			Success:   false,
			Message:   "VALIDATION_ERROR",
			ErrorCode: fmt.Sprintf("BAD_REQUEST | %d", fiber.StatusBadRequest),
			Details:   validationErrs,
		})
	}

	result, err := h.s.Expense.Insert(c, request)
	if err != nil {
		if errors.Is(err, service.BaseErrInsert) {
			return ctx.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
				Success:   false,
				Message:   "FAILED_INSERT",
				ErrorCode: fmt.Sprintf("BAD_REQUEST | %d", fiber.StatusBadRequest),
				Details:   err.Error(),
			})
		}
		if errors.Is(err, service.BaseErrGet) {
			return ctx.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
				Success:   false,
				Message:   "FAILED_GET",
				ErrorCode: fmt.Sprintf("BAD_REQUEST | %d", fiber.StatusBadRequest),
				Details:   err.Error(),
			})
		}
		if errors.Is(err, service.BaseErrUpdate) {
			return ctx.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
				Success:   false,
				Message:   "FAILED_UPDATE",
				ErrorCode: fmt.Sprintf("BAD_REQUEST | %d", fiber.StatusBadRequest),
				Details:   err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Success:   false,
			Message:   "INTERNAL_SERVER_ERROR",
			ErrorCode: fmt.Sprintf("INTERNAL_SERVER_ERROR | %d", fiber.StatusInternalServerError),
			Details:   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Success: true,
		Message: "SUCCESS_INSERT_EXPENSE",
		Data:    result,
	})
}

func (h *ExpenseHandler) UpdateExpense(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	request := new(request.UpdateExpense)

	if err := ctx.BodyParser(request); err != nil {
		log.Errorf("failed to parse request : %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
			Success:   false,
			Message:   "FAILED_PARSE_REQUES",
			ErrorCode: fmt.Sprintf("BAD_REQUEST | %d", fiber.StatusBadRequest),
			Details:   err,
		})
	}

	result, err := h.s.Expense.Update(ctx.Context(), id, request)
	if err != nil {
		if errors.Is(err, service.BaseErrInsert) {
			return ctx.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
				Success:   false,
				Message:   "FAILED_INSERT",
				ErrorCode: fmt.Sprintf("BAD_REQUEST | %d", fiber.StatusBadRequest),
				Details:   err.Error(),
			})
		}
		if errors.Is(err, service.BaseErrGet) {
			return ctx.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
				Success:   false,
				Message:   "FAILED_GET",
				ErrorCode: fmt.Sprintf("BAD_REQUEST | %d", fiber.StatusBadRequest),
				Details:   err.Error(),
			})
		}
		if errors.Is(err, service.BaseErrUpdate) {
			return ctx.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
				Success:   false,
				Message:   "FAILED_UPDATE",
				ErrorCode: fmt.Sprintf("BAD_REQUEST | %d", fiber.StatusBadRequest),
				Details:   err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Success:   false,
			Message:   "INTERNAL_SERVER_ERROR",
			ErrorCode: fmt.Sprintf("INTERNAL_SERVER_ERROR | %d", fiber.StatusInternalServerError),
			Details:   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Success: true,
		Message: "SUCCESS_UPDATE_EXPENSE",
		Data:    result,
	})
}

func (h *ExpenseHandler) GetExpenses(ctx *fiber.Ctx) error {
	year := ctx.QueryInt("year", 0)
	month := ctx.QueryInt("month", 0)

	c := ctx.UserContext()

	result, err := h.s.Expense.GetExpenses(c, year, month)
	if err != nil {
		if errors.Is(err, service.ErrEmpty) {
			return ctx.Status(fiber.StatusNotFound).JSON(utils.ApiResponse{
				Success:   false,
				Message:   "NOT_FOUND",
				ErrorCode: fmt.Sprintf("NOT_FOUND | %d", fiber.StatusNotFound),
				Details:   err.Error(),
			})
		}
		if errors.Is(err, service.ErrInvalid) {
			return ctx.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
				Success:   false,
				Message:   "MISSING_QUERY",
				ErrorCode: fmt.Sprintf("BAD_REQUEST | %d", fiber.StatusBadRequest),
				Details:   err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Success:   false,
			Message:   "INTERNAL_SERVER_ERROR",
			ErrorCode: fmt.Sprintf("INTERNAL_SERVER_ERROR | %d", fiber.StatusInternalServerError),
			Details:   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Success: true,
		Message: "SUCCESS_UPDATE_INCOME",
		Data:    result,
	})
}
