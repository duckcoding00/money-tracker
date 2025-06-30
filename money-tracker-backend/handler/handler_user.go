package handler

import (
	"context"
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

	_, err := h.s.User.Create(ctx.Context(), request)
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

	// send otp for verify
	if err := h.s.Token.ValidationToken(ctx.Context(), request.Username); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
			Success:   false,
			Message:   "BAD_REQUEST",
			ErrorCode: fmt.Sprintf("BAD_REQUEST | %d", fiber.StatusBadRequest),
			Details:   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(utils.ApiResponse{
		Success: true,
		Message: "SUCCESS_CREATED",
		Data:    "success created, please check email for verify user",
	})
}

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
				Details:   err.Error(),
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

func (h *UserHandler) VerifyUser(ctx *fiber.Ctx) error {
	token := ctx.Query("token")

	if token == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
			Success:   false,
			Message:   "MISSING_PARAMETERS",
			ErrorCode: fmt.Sprintf("BAD_REQUEST | %d", fiber.StatusBadRequest),
			Details:   "token and username are required",
		})
	}

	c := context.WithValue(ctx.Context(), "token", token)
	ctx.SetUserContext(c)

	if err := h.s.User.VerifyUser(c); err != nil {
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
			Message:   "INTERNAL_SERVER",
			ErrorCode: fmt.Sprintf("INTERNAL_SERVER_ERROR | %d", fiber.StatusInternalServerError),
			Details:   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Success: true,
		Message: "SUCCESS_VERIFY_USER",
		Data:    "please login again",
	})
}

func (h *UserHandler) ResetPassword(ctx *fiber.Ctx) error {
	request := new(request.NewPassword)

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

	token := ctx.Query("token")
	username := ctx.Query("username")

	if token == "" || username == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
			Success:   false,
			Message:   "MISSING_PARAMETERS",
			ErrorCode: fmt.Sprintf("BAD_REQUEST | %d", fiber.StatusBadRequest),
			Details:   "token and username are required",
		})
	}

	log.Info(token)
	log.Info(username)
	c := context.WithValue(ctx.Context(), "username", username)
	c = context.WithValue(c, "token", token)
	ctx.SetUserContext(c)

	if err := h.s.User.NewPassword(c, request.Password); err != nil {
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
			Message:   "INTERNAL_SERVER",
			ErrorCode: fmt.Sprintf("INTERNAL_SERVER_ERROR | %d", fiber.StatusInternalServerError),
			Details:   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Success: true,
		Message: "SUCCESS_UPDATE_PASSWORD",
		Data:    "please login again",
	})
}

func (h *UserHandler) Profile(ctx *fiber.Ctx) error {
	c := ctx.UserContext()

	result, err := h.s.User.Profile(c)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Success:   false,
			Message:   "INTERNAL_SERVER",
			ErrorCode: fmt.Sprintf("INTERNAL_SERVER_ERROR | %d", fiber.StatusInternalServerError),
			Details:   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Success: true,
		Message: "SUCCESS_UPDATE_PASSWORD",
		Data:    result,
	})
}
