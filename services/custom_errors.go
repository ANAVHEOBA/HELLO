package services

import "errors"

// Define custom error types

var (
	ErrUserNotFound       = errors.New("user does not exist")
	ErrEmailAlreadyExists  = errors.New("email is already registered")
	ErrInvalidCredentials  = errors.New("invalid email or password")
	ErrFailedToCreateUser  = errors.New("failed to create user")
	ErrUnauthorizedAccess  = errors.New("unauthorized access")
)

// CustomError wraps an error with additional context
type CustomError struct {
	Message string
	Err     error
}

// Error implements the error interface for CustomError
func (e *CustomError) Error() string {
	return e.Message + ": " + e.Err.Error()
}

// NewCustomError creates a new CustomError
func NewCustomError(message string, err error) *CustomError {
	return &CustomError{
		Message: message,
		Err:     err,
	}
}
