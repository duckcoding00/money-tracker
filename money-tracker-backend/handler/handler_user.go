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

type UserHandler struct {
	s *service.Service
}

// register handler
func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	request := new(request.UserRequest)

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

	result, err := h.s.User.Create(ctx.Context(), request)
	if err != nil {
		log.Errorf("failed register new user : %v", err)
		if errorhandler.IsDuplicateError(err) {
			return ctx.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
				Success:   false,
				Message:   "DUPLICATE_DATA",
				ErrorCode: fmt.Sprintf("BAD_REQUEST | %d", fiber.StatusBadRequest),
				Details:   err.Error(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Success:   false,
			Message:   "INTERNAL_SERVER",
			ErrorCode: fmt.Sprintf("INTERNAL_SERVER_ERROR | %d", fiber.StatusInternalServerError),
			Details:   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(utils.ApiResponse{
		Success: true,
		Message: "SUCCESS_CREATED",
		Data:    result,
	})
}

// register login
func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	request := new(request.LoginRequest)

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

	result, err := h.s.User.Login(ctx.Context(), request)
	if err != nil {
		log.Errorf("failed to login user : %v", err)
		if errors.Is(err, service.ErrCredentials) {
			return ctx.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
				Success:   false,
				Message:   "BAD_REQUEST",
				ErrorCode: fmt.Sprintf("BAD_REQUEST | %d", fiber.StatusBadRequest),
				Details:   err,
			})
		}

		if errors.Is(err, service.ErrAccessToken) || errors.Is(err, service.ErrRefreshToken) {
			return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
				Success:   false,
				Message:   "TOKEN_GENERATION_FAILED",
				ErrorCode: fmt.Sprintf("INTERNAL_SERVER_ERROR | %d", fiber.StatusInternalServerError),
				Details:   err.Error(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Success:   false,
			Message:   "INTERNAL_SERVER",
			ErrorCode: fmt.Sprintf("INTERNAL_SERVER_ERROR | %d", fiber.StatusInternalServerError),
			Details:   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Success: true,
		Message: "SUCCESS_CREATED",
		Data:    result,
	})
}
