package labelmodel

import "github.com/pkg/errors"

var (
	ErrExistedLabel        = errors.New("label already existed")
	ErrLabelDeleted        = errors.New("label not found")
	ErrLabelNotExistInCard = errors.New("label not exist in card")
)
