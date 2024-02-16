package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// custom errors for various scenarios
var (
	ErrMissingTeacherEmail      = errors.New("no teacher email provided")
	ErrMissingStudentEmail      = errors.New("no student emails provided")
	ErrStudentAlreadyRegistered = errors.New("student already registered")
	ErrTeacherNotFound          = errors.New("teacher not found")
	ErrStudentNotFound          = errors.New("student not found")
	ErrInvalidRequestBody       = errors.New("invalid request body")
	ErrInternalServer           = errors.New("internal server error")
	ErrNoResults                = errors.New("non-existent data found")
)

// NewAPIError creates a new API error response
func NewAPIError(err error) gin.H {
	return gin.H{"message": err.Error()}
}

// ConvertToAPIError Convert database and other known errors to custom API errors
func ConvertToAPIError(err error) error {

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return ErrStudentAlreadyRegistered
		}
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrNoResults
	}
	return ErrInternalServer
}
