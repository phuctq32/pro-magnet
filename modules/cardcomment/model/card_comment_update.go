package cardcommentmodel

import "time"

type CardCommentUpdate struct {
	Content   string    `json:"content" validate:"required" bson:"content"`
	UpdatedAt time.Time `bson:"updatedAt"`
}
