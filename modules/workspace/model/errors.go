package wsmodel

import "github.com/pkg/errors"

var (
	ErrExistedWorkspace = errors.New("workspace name already existed")
)
