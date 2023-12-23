package cardchecklistmodel

type ChecklistItemUpdate struct {
	Title  *string `json:"title,omitempty" bson:"title,omitempty" validate:"omitempty,required"`
	IsDone *bool   `json:"isDone,omitempty" bson:"isDone,omitempty" validate:"omitempty,required"`
}
