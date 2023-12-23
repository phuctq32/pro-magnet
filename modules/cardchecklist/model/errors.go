package cardchecklistmodel

import "github.com/pkg/errors"

var (
	ErrChecklistNotFound     = errors.New("checklist not found")
	ErrChecklistItemNotFound = errors.New("checklist item not found")
)
