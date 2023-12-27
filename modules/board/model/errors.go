package boardmodel

import (
	"errors"
	"pro-magnet/common"
)

var (
	ErrBoardNotFound          = errors.New("board not found")
	ErrBoardDeleted           = errors.New("board not found")
	ErrUserNotBoardAdmin      = errors.New("user not board admin ")
	ErrExistedBoard           = common.NewExistedErr("board name")
	ErrIsNotMemberOfWorkspace = common.NewBadRequestErr(errors.New("user is not workspace member"))
)
