package cardmodel

type CardUpdate struct {
	Title       *string `json:"title,omitempty" validate:"omitempty,required" bson:"title,omitempty"`
	Description *string `json:"description,omitempty" validate:"omitempty,required" bson:"description,omitempty"`
	Cover       *string `json:"cover,omitempty" validate:"omitempty,url" bson:"cover,omitempty"`
}
