package cardchecklistmodel

import "go.mongodb.org/mongo-driver/bson/primitive"

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
