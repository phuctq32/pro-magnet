package authmodel

import "github.com/pkg/errors"

var (
	ErrUserExisted = errors.New("user already existed")
)
