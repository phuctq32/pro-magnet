package cardmodel

import (
	"time"
)

type CardUpdate struct {
	Title       *string `json:"title,omitempty" validate:"omitempty,required" bson:"title,omitempty"`
	Description *string `json:"description,omitempty" validate:"omitempty,required" bson:"description,omitempty"`
	Cover       *string `json:"cover,omitempty" validate:"omitempty,url" bson:"cover,omitempty"`
	IsDone      *bool   `json:"isDone,omitempty" validate:"omitempty" bson:"isDone,omitempty"`
}

type CardDateUpdate struct {
	StartDate *time.Time `json:"startDate,omitempty" validate:"omitempty" bson:"startDate,omitempty"`
	EndDate   *time.Time `json:"endDate,omitempty" validate:"omitempty" bson:"endDate,omitempty"`
}

type CardLabelUpdate struct {
	LabelIds []string `json:""`
}
