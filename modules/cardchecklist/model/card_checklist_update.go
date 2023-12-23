package cardchecklistmodel

type CardChecklistUpdate struct {
	Name string `json:"name" bson:"name" validate:"required"`
}
