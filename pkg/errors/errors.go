package errors

import "fmt"

// ErrorCode represents application error codes
type ErrorCode string

const (
	// Validation errors
	ErrCodeValidation   ErrorCode = "VALIDATION_ERROR"
	ErrCodeInvalidInput ErrorCode = "INVALID_INPUT"

	// Not found errors
	ErrCodeNotFound       ErrorCode = "NOT_FOUND"
	ErrCodePersonNotFound ErrorCode = "PERSON_NOT_FOUND"

	// Conflict errors
	ErrCodeConflict       ErrorCode = "CONFLICT"
	ErrCodeDuplicateEntry ErrorCode = "DUPLICATE_ENTRY"

	// Internal errors
	ErrCodeInternal ErrorCode = "INTERNAL_ERROR"
	ErrCodeDatabase ErrorCode = "DATABASE_ERROR"
	ErrCodeUnknown  ErrorCode = "UNKNOWN_ERROR"

	// Business logic errors
	ErrCodeBusinessRule ErrorCode = "BUSINESS_RULE_VIOLATION"
	ErrCodeConcurrency  ErrorCode = "CONCURRENCY_ERROR"
)

// AppError represents a custom application error
type AppError struct {
	Code    ErrorCode
	Message string
	Err     error
	Details map[string]interface{}
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Unwrap returns the wrapped error
func (e *AppError) Unwrap() error {
	return e.Err
}

// NewValidationError creates a new validation error
func NewValidationError(message string, details map[string]interface{}) *AppError {
	return &AppError{
		Code:    ErrCodeValidation,
		Message: message,
		Details: details,
	}
}

// NewNotFoundError creates a new not found error
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:    ErrCodeNotFound,
		Message: message,
	}
}

// NewInternalError creates a new internal error
func NewInternalError(message string, err error) *AppError {
	return &AppError{
		Code:    ErrCodeInternal,
		Message: message,
		Err:     err,
	}
}

// NewDatabaseError creates a new database error
func NewDatabaseError(message string, err error) *AppError {
	return &AppError{
		Code:    ErrCodeDatabase,
		Message: message,
		Err:     err,
	}
}

// NewConflictError creates a new conflict error
func NewConflictError(message string) *AppError {
	return &AppError{
		Code:    ErrCodeConflict,
		Message: message,
	}
}

// NewConcurrencyError creates a new concurrency error
func NewConcurrencyError(message string) *AppError {
	return &AppError{
		Code:    ErrCodeConcurrency,
		Message: message,
	}
}
