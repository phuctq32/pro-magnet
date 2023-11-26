package columnmodel

type ColumnStatus uint8

type Column struct {
	Id             *string      `json:"_id,omitempty" bson:"_id,omitempty"`
	Status         ColumnStatus `json:"-" bson:"status"`
	Title          string       `json:"title" bson:"title"`
	BoardId        string       `json:"boardId" bson:"boardId"`
	OrderedCardIds []string     `json:"orderedCardIds" bson:"orderedCardIds"`
}
