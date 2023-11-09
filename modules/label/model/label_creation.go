package labelmodel

type LabelCreation struct {
	Title   string  `json:"title" validate:"required"`
	Color   string  `json:"color" validate:"required,hexcolor"`
	BoardId string  `json:"boardId" validate:"required,mongodb"`
	CardId  *string `json:"cardId,omitempty" validate:"omitempty,mongodb"`
}
