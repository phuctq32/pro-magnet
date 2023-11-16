package carduc

import (
	"context"
	"pro-magnet/common"
	cardmodel "pro-magnet/modules/card/model"
)

func (uc *cardUseCase) GetCardById(ctx context.Context, id string) (*cardmodel.Card, error) {
	card, err := uc.cardRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	if card.Status == cardmodel.Deleted {
		return nil, common.NewBadRequestErr(cardmodel.ErrCardDeleted)
	}

	// Aggregate data
	if err = uc.cardAgg.Aggregate(ctx, card); err != nil {
		return nil, err
	}

	return card, nil
}
