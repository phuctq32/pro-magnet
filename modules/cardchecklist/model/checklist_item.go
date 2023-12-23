package cardchecklistmodel

import "go.mongodb.org/mongo-driver/bson/primitive"

type ChecklistItem struct {
	Id     *string `json:"_id,omitempty" bson:"_id,omitempty"`
	Title  string  `json:"title" bson:"title" validate:"required"`
	IsDone bool    `json:"isDone" bson:"isDone"`
}

type ChecklistItemInsert struct {
	Id     primitive.ObjectID `bson:"_id"`
	Title  string             `bson:"title"`
	IsDone bool               `bson:"isDone"`
}
