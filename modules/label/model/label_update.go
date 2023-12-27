package labelmodel

type LabelUpdate struct {
	Title *string `json:"title" validate:"required"`
	Color *string `json:"color" validate:"required,hexcolor"`
}
