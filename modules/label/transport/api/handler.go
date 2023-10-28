package labelapi

import (
	"context"
	labelmodel "pro-magnet/modules/label/model"
)

type LabelUseCase interface {
	CreateLabel(ctx context.Context, data *labelmodel.LabelCreation) (*labelmodel.Label, error)
}

type labelHandler struct {
	uc LabelUseCase
}

func NewLabelHandler(uc LabelUseCase) *labelHandler {
	return &labelHandler{uc: uc}
}
