package labeluc

import (
	"context"
	"github.com/pkg/errors"
	"pro-magnet/common"
	labelmodel "pro-magnet/modules/label/model"
)

func (uc *labelUseCase) CreateLabel(
	ctx context.Context,
	data *labelmodel.LabelCreation,
) (*labelmodel.Label, error) {
	// Check label valid:
	// - Label has not existed in board (don't match both title and color)
	isExisted, err := uc.labelRepo.ExistsInBoard(ctx, data)
	if err != nil {
		return nil, err
	}
	if isExisted {
		return nil, common.NewBadRequestErr(labelmodel.ErrExistedLabel)
	}

	// Should be in a transaction
	var label *labelmodel.Label
	var e error

	err = uc.labelRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		label, e = uc.labelRepo.Create(txCtx, data)
		if e != nil {
			return e
		}

		// Save label to card if label was created in card
		if data.CardId != nil {

		} else {
			return common.NewBadRequestErr(errors.New("err tx"))
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return label, nil
}
