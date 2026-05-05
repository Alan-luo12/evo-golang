package errors

import (
	"net/http"
)

type ErrorType int

const (
	ErrorTypeUser ErrorType = iota
	ErrorTypeSystem
)

type AppError struct {
	Type ErrorType
	Msg  string
	Code int
	Err  error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Msg + ":" + e.Err.Error()
	}

	return e.Msg
}

func (e *AppError) HTTPStatus() int {
	if e.Type == ErrorTypeSystem {
		return http.StatusInternalServerError
	}

	return http.StatusBadRequest
}

func NewUserError(code int, msg string, err error) *AppError {
	return &AppError{
		Type: ErrorTypeUser,
		Code: code,
		Msg:  msg,
		Err:  err,
	}
}

func NewSystemError(code int, msg string, err error) *AppError {
	return &AppError{
		Type: ErrorTypeSystem,
		Code: code,
		Msg:  msg,
		Err:  err,
	}
}
