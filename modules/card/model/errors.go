package cardmodel

import "github.com/pkg/errors"

var (
	ErrCardNotFound          = errors.New("card not found")
	ErrCardDeleted           = errors.New("card already deleted")
	ErrUserAddedToCardBefore = errors.New("user already added to card")
	ErrUserNotExistInCard    = errors.New("user not a member of card")
)
