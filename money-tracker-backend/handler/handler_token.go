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

type TokenHandler struct {
	s *service.Service
}

func (h *TokenHandler) RefreshToken(ctx *fiber.Ctx) error {

	result, err := h.s.Token.GenerateAccessToken(ctx.UserContext())
	if err != nil {
		if errors.Is(err, service.ExpiredRefreshToken) {
			return ctx.Status(fiber.StatusUnauthorized).JSON(utils.ApiResponse{
				Success:   false,
				Message:   "REFRESH_TOKEN_EXPIRED",
				ErrorCode: fmt.Sprintf("UNAUTHORIZED | %d", fiber.StatusUnauthorized),
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
		Message: "success generated access token",
		Data:    result,
	})
}

func (h *TokenHandler) ResetToken(ctx *fiber.Ctx) error {

	request := new(request.ResetUser)

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

	if err := h.s.Token.ResetToken(ctx.Context(), request.Username); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(utils.ApiResponse{
			Success: true,
			Message: "send token if users exists",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Success: true,
		Message: "send token if users exists",
	})
}

func (h *TokenHandler) VerifyToken(ctx *fiber.Ctx) error {
	request := new(request.VerifyToken)

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

	result, err := h.s.Token.VerifyResetToken(ctx.Context(), request)
	if err != nil {
		if errors.Is(err, service.InvalidToken) {
			return ctx.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
				Success:   false,
				Message:   "OTP_TOKEN_INVALID",
				ErrorCode: fmt.Sprintf("BAD_REQUEST | %d", fiber.StatusBadRequest),
				Details:   err.Error(),
			})
		}
		if errors.Is(err, service.ExpiredToken) {
			return ctx.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
				Success:   false,
				Message:   "OTP_TOKEN_EXPIRED",
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
		Message: "success generated access token",
		Data:    result,
	})
}
