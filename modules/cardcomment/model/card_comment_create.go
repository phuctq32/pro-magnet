package cardcommentmodel

type CardCommentCreate struct {
	Content string `json:"content" validate:"required"`
	UserId  string
}
