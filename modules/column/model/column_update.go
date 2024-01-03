package columnmodel

import "go.mongodb.org/mongo-driver/bson/primitive"

type ColumnUpdate struct {
	Title           *string              `json:"title,omitempty" bson:"title,omitempty" validate:"omitempty,required"`
	OrderedCardIds  []string             `json:"orderedCardIds,omitempty" bson:"-" validate:"omitempty,dive,required"`
	OrderedCardOids []primitive.ObjectID `json:"-" bson:"orderedCardIds,omitempty"`
}

func (cu *ColumnUpdate) ToUpdateData() *ColumnUpdate {
	if cu.OrderedCardIds != nil {
		cu.OrderedCardOids = make([]primitive.ObjectID, len(cu.OrderedCardIds))
		for i := 0; i < len(cu.OrderedCardOids); i++ {
			cu.OrderedCardOids[i], _ = primitive.ObjectIDFromHex(cu.OrderedCardIds[i])
		}
	}

	return cu
}
