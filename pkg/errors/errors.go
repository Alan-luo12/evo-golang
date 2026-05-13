package errors

import (
	"net/http"
)

// 错误类型
type ErrorType int

const (
	//用户错误
	ErrorTypeUser ErrorType = iota
	//系统错误
	ErrorTypeSystem

	ErrorTypeLimitExceeded

	ErrorTypeNotFound

	ErrorTypeUnauthorized

	ErrorTypeConflict
)

// AppError结构体
type AppError struct {
	Type ErrorType
	Msg  string
	Code int
	Err  error
}

// 错误信息
func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Msg + ":" + e.Err.Error()
	}

	return e.Msg
}

// HTTP状态码
func (e *AppError) HTTPStatus() int {
	if e.Type == ErrorTypeSystem {
		return http.StatusInternalServerError
	}
	//如果错误类型是限流错误，则返回429
	if e.Type == ErrorTypeLimitExceeded {
		return http.StatusTooManyRequests
	}
	//如果错误类型是未找到错误，则返回404
	if e.Type == ErrorTypeNotFound {
		return http.StatusNotFound
	}
	//如果错误类型是未授权错误，则返回401
	if e.Type == ErrorTypeUnauthorized {
		return http.StatusUnauthorized
	}

	if e.Type == ErrorTypeConflict {
		return http.StatusConflict
	}
	//默认返回400
	return http.StatusBadRequest
}

// 创建用户错误
func NewUserError(code int, msg string, err error) *AppError {
	return &AppError{
		Type: ErrorTypeUser,
		Code: code,
		Msg:  msg,
		Err:  err,
	}
}

// 创建系统错误
func NewSystemError(code int, msg string, err error) *AppError {
	return &AppError{
		Type: ErrorTypeSystem,
		Code: code,
		Msg:  msg,
		Err:  err,
	}
}

// 创建限流错误
func NewLimitExceededError(code int, msg string, err error) *AppError {
	return &AppError{
		Type: ErrorTypeLimitExceeded,
		Code: code,
		Msg:  msg,
		Err:  err,
	}
}

// 创建未找到错误
func NewNotFoundError(code int, msg string, err error) *AppError {
	return &AppError{
		Type: ErrorTypeNotFound,
		Code: code,
		Msg:  msg,
		Err:  err,
	}
}

// 创建未授权错误
func NewUnauthorizedError(code int, msg string, err error) *AppError {
	return &AppError{
		Type: ErrorTypeUnauthorized,
		Code: code,
		Msg:  msg,
		Err:  err,
	}
}

func NewConflictError(code int, msg string, err error) *AppError {
	return &AppError{
		Type: ErrorTypeConflict,
		Code: code,
		Msg:  msg,
		Err:  err,
	}
}
