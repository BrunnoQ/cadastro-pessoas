package errors

import "fmt"

// AppError represents a custom application error
type AppError struct {
	Code    ErrorCode              `json:"code"`
	Message string                 `json:"message"`
	Err     error                  `json:"-"`
	Details map[string]interface{} `json:"details,omitempty"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (caused by: %v)", e.Code, e.Message, e.Err)
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
		Code:    CodeValidation,
		Message: message,
		Details: details,
	}
}

// NewNotFoundError creates a new not found error
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:    CodeNotFound,
		Message: message,
	}
}

// NewInternalError creates a new internal error
func NewInternalError(message string, err error) *AppError {
	return &AppError{
		Code:    CodeInternal,
		Message: message,
		Err:     err,
	}
}

// NewDatabaseError creates a new database error
func NewDatabaseError(message string, err error) *AppError {
	return &AppError{
		Code:    CodeDatabase,
		Message: message,
		Err:     err,
	}
}

// NewConflictError creates a new conflict error
func NewConflictError(message string) *AppError {
	return &AppError{
		Code:    CodeConflict,
		Message: message,
	}
}

// NewConcurrencyError creates a new concurrency error
func NewConcurrencyError(message string) *AppError {
	return &AppError{
		Code:    CodeConcurrency,
		Message: message,
	}
}

// IsNotFoundError checks if an error is a not found error
func IsNotFoundError(err error) bool {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Code == CodeNotFound || appErr.Code == CodePersonNotFound
	}
	return false
}

// IsConcurrencyError checks if an error is a concurrency error
func IsConcurrencyError(err error) bool {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Code == CodeConcurrency
	}
	return false
}
