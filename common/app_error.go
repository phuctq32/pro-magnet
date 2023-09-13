package common

import (
	"fmt"
	"net/http"
)

const (
	InternalErrKey     string = "INTERNAL_SERVER_ERROR"
	BadRequestErrKey          = "BAD_REQUEST"
	NotFoundErrKey            = "NOT_FOUND"
	NoPermissionErrKey        = "NO_PERMISSION"
)

type AppError struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"status_code"`
	Key        string `json:"error_key"`
	Message    string `json:"message"`
	Log        string `json:"-"`
	Err        error  `json:"-"`
	// ValidationErrs
}

func (e *AppError) Error() string {
	return e.RootErr().Error()
}

func (e *AppError) RootErr() error {
	if err, ok := e.Err.(*AppError); ok {
		return err.RootErr()
	}
	return e.Err
}

func NewErrResponse(statusCode int, key, msg, log string, err error) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Key:        key,
		Message:    msg,
		Log:        log,
		Err:        err,
	}
}

func NewServerErr(err error) *AppError {
	return NewErrResponse(
		http.StatusInternalServerError,
		InternalErrKey,
		"Internal Server Error",
		err.Error(),
		err,
	)
}

func NewBadRequestErr(err error, msg ...string) *AppError {
	message := err.Error()
	if len(msg) > 0 {
		message = msg[0]
	}

	return NewErrResponse(
		http.StatusBadRequest,
		BadRequestErrKey,
		message,
		err.Error(),
		err,
	)
}

func NewNotFoundErr(entity string, err error) *AppError {
	return NewErrResponse(
		http.StatusNotFound,
		NotFoundErrKey,
		fmt.Sprintf("%v not found", entity),
		err.Error(),
		err,
	)
}

func NewNoPermissionErr(err error) *AppError {
	return NewErrResponse(
		http.StatusUnprocessableEntity,
		NoPermissionErrKey,
		"access denied",
		err.Error(),
		err,
	)
}
