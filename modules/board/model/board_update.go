package boardmodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type BoardUpdate struct {
	Name              *string              `json:"name,omitempty" validate:"omitempty,required" bson:"name,omitempty"`
	OrderedColumnIds  []string             `json:"orderedColumnIds,omitempty" validate:"omitempty,dive,required,mongodb" bson:"-"`
	OrderedColumnOids []primitive.ObjectID `json:"-" bson:"orderedColumnIds,omitempty"`
	UpdatedAt         time.Time
}

func (bu *BoardUpdate) ToUpdateData() *BoardUpdate {
	if bu.OrderedColumnIds != nil {
		bu.OrderedColumnOids = make([]primitive.ObjectID, len(bu.OrderedColumnIds))
		for i := 0; i < len(bu.OrderedColumnIds); i++ {
			bu.OrderedColumnOids[i], _ = primitive.ObjectIDFromHex(bu.OrderedColumnIds[i])
		}
	}

	return bu
}
