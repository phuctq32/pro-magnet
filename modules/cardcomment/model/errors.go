package cardcommentmodel

import "github.com/pkg/errors"

var (
	ErrCommentNotFound  = errors.New("comment not found")
	ErrNotCommentAuthor = errors.New("user is not comment's author")
)
