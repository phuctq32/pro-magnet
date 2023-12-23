package boardmodel

import (
	"errors"
	"pro-magnet/common"
)

var (
	ErrBoardNotFound          = errors.New("board not found")
	ErrExistedBoard           = common.NewExistedErr("board")
	ErrIsNotMemberOfWorkspace = common.NewBadRequestErr(errors.New("user is not workspace member"))
)
