package common

import (
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"pro-magnet/components/validator"
)

const (
	InternalErrKey     string = "INTERNAL_SERVER_ERROR"
	BadRequestErrKey          = "BAD_REQUEST"
	NotFoundErrKey            = "NOT_FOUND"
	UnauthorizedErrKey        = "UNAUTHORIZED"
	NoPermissionErrKey        = "NO_PERMISSION"
	ExistedErrKey             = "EXISTED"
	ValidationErrKey          = "VALIDATION_FAILED"
)

var (
	ErrNotFound = mongo.ErrNoDocuments.Error()
)

type AppError struct {
	Success        bool                        `json:"success"`
	StatusCode     int                         `json:"statusCode"`
	Key            string                      `json:"errorKey"`
	Message        string                      `json:"message"`
	Log            string                      `json:"-"`
	Err            error                       `json:"-"`
	ValidationErrs []validator.ValidationError `json:"validationErrors,omitempty"`
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

func NewExistedErr(entity string) *AppError {
	return NewErrResponse(
		http.StatusUnprocessableEntity,
		ExistedErrKey,
		fmt.Sprintf("%v already existed", entity),
		fmt.Sprintf("%v already existed", entity),
		errors.New(fmt.Sprintf("%v already existed", entity)),
	)
}

func NewNoPermissionErr(err error) *AppError {
	return NewErrResponse(
		http.StatusForbidden,
		NoPermissionErrKey,
		"access denied",
		err.Error(),
		err,
	)
}

func NewUnauthorizedErr(err error) *AppError {
	return NewErrResponse(
		http.StatusUnauthorized,
		UnauthorizedErrKey,
		err.Error(),
		err.Error(),
		err,
	)
}

func NewValidationErrors(errs []validator.ValidationError) *AppError {
	return &AppError{
		StatusCode:     http.StatusUnprocessableEntity,
		Key:            ValidationErrKey,
		Message:        fmt.Sprintf("field: '%s' error: %s", errs[0].Field, errs[0].Message),
		Err:            &errs[0],
		ValidationErrs: errs,
	}
}
