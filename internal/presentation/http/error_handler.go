package http

import (
	"net/http"

	"github.com/BrunnoQ/cadastro-pessoas/pkg/errors"
)

// ErrorCode maps internal error codes to HTTP status codes
var errorCodeToHTTPStatus = map[errors.ErrorCode]int{
	// Validation errors -> 400 Bad Request
	errors.CodeValidation:   http.StatusBadRequest,
	errors.CodeInvalidInput: http.StatusBadRequest,
	errors.CodeInvalidEmail: http.StatusBadRequest,
	errors.CodeInvalidPhone: http.StatusBadRequest,
	errors.CodeInvalidDate:  http.StatusBadRequest,
	errors.CodeInvalidName:  http.StatusBadRequest,
	errors.CodeInvalidSex:   http.StatusBadRequest,
	errors.CodeInvalidAge:   http.StatusBadRequest,
	errors.CodeBadRequest:   http.StatusBadRequest,

	// Not found errors -> 404 Not Found
	errors.CodeNotFound:       http.StatusNotFound,
	errors.CodePersonNotFound: http.StatusNotFound,

	// Conflict errors -> 409 Conflict
	errors.CodeConflict:    http.StatusConflict,
	errors.CodeDuplicate:   http.StatusConflict,
	errors.CodeConcurrency: http.StatusConflict,

	// Authorization errors
	errors.CodeUnauthorized: http.StatusUnauthorized,
	errors.CodeForbidden:    http.StatusForbidden,

	// Rate limiting
	errors.CodeTooManyRequests: http.StatusTooManyRequests,

	// Business rule errors -> 422 Unprocessable Entity
	errors.CodeBusinessRule: http.StatusUnprocessableEntity,

	// Internal errors -> 500 Internal Server Error
	errors.CodeInternal: http.StatusInternalServerError,
	errors.CodeDatabase: http.StatusInternalServerError,
	errors.CodeUnknown:  http.StatusInternalServerError,
}

// ErrorResponse represents a standardized error response
type ErrorResponse struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
}

// ToErrorResponse converts an AppError to HTTP error response
func ToErrorResponse(err error) (int, ErrorResponse) {
	// Try to cast to AppError
	appErr, ok := err.(*errors.AppError)
	if !ok {
		// Unknown error, return generic internal error
		return http.StatusInternalServerError, ErrorResponse{
			Code:    string(errors.CodeInternal),
			Message: "An internal error occurred",
		}
	}

	// Map error code to HTTP status
	statusCode, exists := errorCodeToHTTPStatus[appErr.Code]
	if !exists {
		statusCode = http.StatusInternalServerError
	}

	// Build response
	response := ErrorResponse{
		Code:    string(appErr.Code),
		Message: appErr.Message,
		Details: appErr.Details,
	}

	return statusCode, response
}

// GetHTTPStatus returns HTTP status code for error code
func GetHTTPStatus(code errors.ErrorCode) int {
	if status, exists := errorCodeToHTTPStatus[code]; exists {
		return status
	}
	return http.StatusInternalServerError
}
