package cardmodel

import (
	usermodel "pro-magnet/modules/user/model"
	"time"
)

type CardComment struct {
	Id        *string   `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
	Content   string    `json:"content" bson:"content"`
	AuthorId  string    `json:"-" bson:"authorId"`

	// Aggregated data
	Author *usermodel.User `json:"author" bson:"-"`
}
