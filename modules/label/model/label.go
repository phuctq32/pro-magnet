package labelmodel

import "go.mongodb.org/mongo-driver/bson/primitive"

type LabelInsert struct {
	Status  LabelStatus        `bson:"status"`
	Title   string             `bson:"title"`
	Color   string             `bson:"color"`
	BoardId primitive.ObjectID `bson:"boardId"`
}

type Label struct {
	Id      *string     `json:"_id,omitempty" bson:"_id,omitempty"`
	Status  LabelStatus `json:"-" bson:"status"`
	Title   string      `json:"title" bson:"title"`
	Color   string      `json:"color" bson:"color"`
	BoardId string      `json:"-" bson:"boardId"`
}
