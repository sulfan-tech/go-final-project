package customerror

import "fmt"

// CustomError is a custom error type that implements the error interface.
type CustomError struct {
	Code    int    // Custom error code
	Message string // Error message
}

// Error returns the error message of the CustomError.
func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// NewError creates a new CustomError with the given error code and message.
func NewError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}
