package boardmodel

import "time"

type BoardUpdate struct {
	Name             *string   `json:"name,omitempty" validate:"omitempty,required" bson:"name,omitempty"`
	OrderedColumnIds *[]string `json:"orderedColumnIds,omitempty" validate:"omitempty,dive,required,mongodb" bson:"orderedColumnIds,omitempty"`
	UpdatedAt        time.Time
}
