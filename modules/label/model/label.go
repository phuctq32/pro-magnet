package labelmodel

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	LabelCollectionName = "labels"
)

type LabelInsert struct {
	Title   string             `json:"title" bson:"title"`
	Color   string             `json:"color" bson:"color"`
	BoardId primitive.ObjectID `json:"-" bson:"boardId"`
}

type Label struct {
	Id      *string `json:"_id,omitempty" bson:"_id,omitempty"`
	Title   string  `json:"title" bson:"title"`
	Color   string  `json:"color" bson:"color"`
	BoardId string  `json:"-" bson:"boardId"`
}
