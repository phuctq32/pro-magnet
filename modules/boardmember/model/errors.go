package bmmodel

import "github.com/pkg/errors"

var (
	ErrUserNotBoardMember       = errors.New("user is not board member")
	ErrUserIsABoardMemberBefore = errors.New("user is a board member before")
)
