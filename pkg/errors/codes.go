package errors

// ErrorCode represents application error codes
type ErrorCode string

const (
	// Validation errors
	CodeValidation   ErrorCode = "VALIDATION_ERROR"
	CodeInvalidInput ErrorCode = "INVALID_INPUT"
	CodeInvalidEmail ErrorCode = "INVALID_EMAIL"
	CodeInvalidPhone ErrorCode = "INVALID_PHONE"
	CodeInvalidDate  ErrorCode = "INVALID_DATE"
	CodeInvalidName  ErrorCode = "INVALID_NAME"
	CodeInvalidSex   ErrorCode = "INVALID_SEX"
	CodeInvalidAge   ErrorCode = "INVALID_AGE"

	// Not found errors
	CodeNotFound       ErrorCode = "NOT_FOUND"
	CodePersonNotFound ErrorCode = "PERSON_NOT_FOUND"

	// Conflict errors
	CodeConflict  ErrorCode = "CONFLICT"
	CodeDuplicate ErrorCode = "DUPLICATE_ENTRY"

	// Internal errors
	CodeInternal ErrorCode = "INTERNAL_ERROR"
	CodeDatabase ErrorCode = "DATABASE_ERROR"
	CodeUnknown  ErrorCode = "UNKNOWN_ERROR"

	// Business logic errors
	CodeBusinessRule ErrorCode = "BUSINESS_RULE_VIOLATION"
	CodeConcurrency  ErrorCode = "CONCURRENCY_ERROR"

	// Request errors
	CodeBadRequest      ErrorCode = "BAD_REQUEST"
	CodeUnauthorized    ErrorCode = "UNAUTHORIZED"
	CodeForbidden       ErrorCode = "FORBIDDEN"
	CodeTooManyRequests ErrorCode = "TOO_MANY_REQUESTS"
)
