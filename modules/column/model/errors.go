package columnmodel

import "github.com/pkg/errors"

var (
	ErrNotBoardMember = errors.New("user is not board member")
	ErrColumnNotFound = errors.New("column not found")
)
