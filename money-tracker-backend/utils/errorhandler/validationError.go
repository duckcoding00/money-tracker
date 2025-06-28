package errorhandler

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidationErrors(errs validator.ValidationErrors) map[string]string {
	errorValidation := make(map[string]string)

	tagMessages := map[string]string{
		"required": "is required",
		"email":    "must be a valid email",
		"min":      "must be at least %s characters",
		"max":      "must be at most %s characters",
		"len":      "must be exactly %s characters",
		"gt":       "must be greater than %s",
		"gte":      "must be greater than or equal to %s",
		"lt":       "must be less than %s",
		"lte":      "must be less than or equal to %s",
		"oneof":    "must be one of: %s",
		"numeric":  "must be numeric",
		"alpha":    "must contain only letters",
		"alphanum": "must contain only letters and numbers",
	}

	for _, err := range errs {
		field := err.Field()
		tag := err.Tag()
		param := err.Param()

		message, exists := tagMessages[tag]
		if !exists {
			message = "is invalid"
		}

		if param != "" && strings.Contains(message, "%s") {
			errorValidation[field] = fmt.Sprintf("%s %s", field, fmt.Sprintf(message, param))
		} else {
			errorValidation[field] = fmt.Sprintf("%s %s", field, message)
		}
	}

	return errorValidation
}
