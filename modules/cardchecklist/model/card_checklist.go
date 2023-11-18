package cardchecklistmodel

import "go.mongodb.org/mongo-driver/bson/primitive"

type ChecklistItem struct {
	Id     *string `json:"_id,omitempty" bson:"_id,omitempty"`
	Title  string  `json:"title" bson:"title"`
	IsDone bool    `json:"isDone" bson:"isDone"`
}

type CardChecklist struct {
	Id    *string         `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string          `json:"name" bson:"name" validate:"required"`
	Items []ChecklistItem `json:"items" bson:"items"`
}

type CardChecklistInsert struct {
	Id    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Items []ChecklistItem    `bson:"items"`
}
