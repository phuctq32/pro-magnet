package labeluc

import (
	"context"
	"pro-magnet/common"
	cardmodel "pro-magnet/modules/card/model"
	labelmodel "pro-magnet/modules/label/model"
)

func (uc *labelUseCase) CreateLabel(
	ctx context.Context,
	data *labelmodel.LabelCreation,
) (*labelmodel.Label, error) {
	if err := uc.validateLabel(ctx, nil, data.BoardId, data.Title, data.Color); err != nil {
		return nil, err
	}

	// Should be in a transaction
	var label *labelmodel.Label
	var e error

	err := uc.labelRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		label, e = uc.labelRepo.Create(txCtx, data)
		if e != nil {
			return e
		}

		// Save label to card if label was created in card
		if data.CardId != nil {
			if e = uc.validateCard(ctx, *data.CardId); e != nil {
				return e
			}

			// Add label to card
			if e = uc.cardRepo.UpdateLabel(ctx, *data.CardId, *label.Id); e != nil {
				return e
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return label, nil
}

func (uc *labelUseCase) validateLabel(ctx context.Context, labelId *string, boardId, title, color string) error {
	// Check label valid:
	// - Label has not existed in board (don't match both title and color)
	isExisted, err := uc.labelRepo.ExistsInBoard(ctx, labelId, boardId, title, color)
	if err != nil {
		return err
	}
	if isExisted {
		return common.NewBadRequestErr(labelmodel.ErrExistedLabel)
	}

	return nil
}

func (uc *labelUseCase) validateCard(ctx context.Context, cardId string) error {
	card, err := uc.cardRepo.FindById(ctx, cardId)
	if err != nil {
		return err
	}
	if card.Status == cardmodel.Deleted {
		return common.NewBadRequestErr(cardmodel.ErrCardDeleted)
	}

	return nil
}
