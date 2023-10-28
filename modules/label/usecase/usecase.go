package labeluc

import (
	"context"
	labelmodel "pro-magnet/modules/label/model"
)

type LabelRepository interface {
	Create(ctx context.Context, data *labelmodel.LabelCreation) (*labelmodel.Label, error)
	ExistsInBoard(ctx context.Context, data *labelmodel.LabelCreation) (bool, error)
	WithinTransaction(ctx context.Context, fn func(context.Context) error) error
}

type labelUseCase struct {
	labelRepo LabelRepository
}

func NewLabelUseCase(
	labelRepo LabelRepository,
) *labelUseCase {
	return &labelUseCase{
		labelRepo: labelRepo,
	}
}
