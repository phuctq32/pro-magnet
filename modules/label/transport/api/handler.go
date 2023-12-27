package labelapi

import (
	"context"
	labelmodel "pro-magnet/modules/label/model"
)

type LabelUseCase interface {
	CreateLabel(ctx context.Context, data *labelmodel.LabelCreation) (*labelmodel.Label, error)
	UpdateLabel(ctx context.Context, requesterId, labelId string, updateData *labelmodel.LabelUpdate) error
}

type labelHandler struct {
	uc LabelUseCase
}

func NewLabelHandler(uc LabelUseCase) *labelHandler {
	return &labelHandler{uc: uc}
}
