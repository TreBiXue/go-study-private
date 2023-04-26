package errorhandler

import "net/http"

type ErrorType string

const (
	ErrUnauthorized    ErrorType = "unauthorized"
	ErrBadParams       ErrorType = "bad_params"
	ErrNotFound        ErrorType = "not_found"
	ErrUniqueViolation ErrorType = "unique_violation"
	ErrDatabase        ErrorType = "database_error"
	ErrUnknown         ErrorType = "unknown_error"
)

var codeStatusMap = map[ErrorType]int{
	ErrUnauthorized:    http.StatusUnauthorized,
	ErrBadParams:       http.StatusBadRequest,
	ErrNotFound:        http.StatusNotFound,
	ErrDatabase:        http.StatusInternalServerError,
	ErrUniqueViolation: http.StatusConflict,
	ErrUnknown:         http.StatusInternalServerError,
}

func (e ErrorType) Error() string {
	return string(e)
}

func GetHTTPStatus(code ErrorType) int {
	return codeStatusMap[code]
}
