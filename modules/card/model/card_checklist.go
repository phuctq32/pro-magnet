package cardmodel

type ChecklistItem struct {
	Id     *string `json:"_id,omitempty" bson:"_id,omitempty"`
	Title  string  `json:"title" bson:"title"`
	IsDone bool    `json:"isDone" bson:"isDone"`
}

type CardChecklist struct {
	Id    *string         `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string          `json:"name" bson:"name"`
	Items []ChecklistItem `json:"items" bson:"items"`
}
