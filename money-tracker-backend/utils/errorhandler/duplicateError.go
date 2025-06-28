package errorhandler

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
)

type DuplicateError struct {
	Field   string
	Value   string
	Message string
}

func (e *DuplicateError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return fmt.Sprintf("%s '%s' already exists", e.Field, e.Value)
}

func HandleDuplicateError(err error, email, username string) error {
	var pgErr *pgconn.PgError

	if !errors.As(err, &pgErr) || pgErr.Code != "23505" {
		return err
	}

	constraint := strings.ToLower(pgErr.ConstraintName)

	switch {
	case strings.Contains(constraint, "email"):
		return &DuplicateError{Field: "email", Value: email, Message: "Email already exists"}
	case strings.Contains(constraint, "username"):
		return &DuplicateError{Field: "username", Value: username, Message: "Username already exists"}
	default:
		return &DuplicateError{Message: "Data already exists"}
	}
}

func IsDuplicateError(err error) bool {
	var duplicateErr *DuplicateError
	return errors.As(err, &duplicateErr)
}
