package usermodel

import "github.com/pkg/errors"

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrUserNotVerified     = errors.New("user not verified")
	ErrIncorrectPassword   = errors.New("password incorrect")
	ErrSkillAlreadyExisted = errors.New("skill already existed")
	ErrSkillNotFound       = errors.New("skill not found")
)
