package wsmodel

import "time"

type WorkspaceUpdate struct {
	Name      *string   `json:"name,omitempty" validate:"omitempty,required" bson:"name,omitempty"`
	Image     *string   `json:"image,omitempty" validate:"omitempty,required,url" bson:"image,omitempty"`
	UpdatedAt time.Time `bson:"updatedAt"`
}
