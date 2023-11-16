package cardmodel

import "github.com/pkg/errors"

var (
	ErrCardNotFound = errors.New("card not found")
	ErrCardDeleted  = errors.New("card already deleted")
)
