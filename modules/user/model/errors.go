package usermodel

import "github.com/pkg/errors"

var (
	ErrUserNotFound = errors.New("user not found")
)
