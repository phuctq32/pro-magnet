package cardmodel

import "time"

type CardAttachment struct {
	Id        *string   `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	BoardId   string    `json:"-" bson:"boardId"`
	CardId    string    `json:"cardId" bson:"cardId"`
	URL       string    `json:"url" bson:"url"`
	FileName  string    `json:"fileName" bson:"fileName"`
	Extension string    `json:"extension" bson:"extension"`
}
