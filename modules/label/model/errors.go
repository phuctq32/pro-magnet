package labelmodel

import "github.com/pkg/errors"

var (
	ErrExistedLabel            = errors.New("label already existed")
	ErrLabelDeleted            = errors.New("label not found")
	ErrLabelAlreadyExistInCard = errors.New("label already existed in card")
	ErrLabelNotExistInCard     = errors.New("label not exist in card")
	ErrLabelNotExistInBoard    = errors.New("label not exist in board")
)
