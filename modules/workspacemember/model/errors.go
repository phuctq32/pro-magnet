package wsmembermodel

import "github.com/pkg/errors"

var (
	ErrUserNotWorkspaceOwner      = errors.New("user is not workspace admin")
	ErrUserAlreadyWorkspaceMember = errors.New("user already a workspace member")
)
