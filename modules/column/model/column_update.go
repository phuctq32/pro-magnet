package columnmodel

type ColumnUpdate struct {
	Title          *string  `json:"title,omitempty" bson:"title,omitempty" validate:"omitempty,required"`
	OrderedCardIds []string `json:"orderedCardIds,omitempty" bson:"orderedCardIds,omitempty" validate:"omitempty,dive,required"`
}
