package boardmodel

import (
	"errors"
	"pro-magnet/common"
)

var (
	ErrExistedBoard           = common.NewExistedErr("board")
	ErrIsNotMemberOfWorkspace = common.NewBadRequestErr(errors.New("user is not workspace member"))
)
