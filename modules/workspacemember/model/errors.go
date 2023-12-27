package wsmembermodel

import "github.com/pkg/errors"

var (
	ErrUserNotWorkspaceOwner      = errors.New("user is not workspace admin")
	ErrUserNotWorkspaceMember     = errors.New("user is not workspace member")
	ErrUserAlreadyWorkspaceMember = errors.New("user already a workspace member")
	ErrCanNotRemoveWorkspaceOwner = errors.New("can not remove workspace admin")
)
