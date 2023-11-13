package carduc

import (
	"context"
	cardmodel "pro-magnet/modules/card/model"
)

func (uc *cardUseCase) UpdateCardById(
	ctx context.Context,
	cardId string,
	data *cardmodel.CardUpdate,
) (*cardmodel.Card, error) {
	// Check user is a member of card's board

	card, err := uc.cardRepo.UpdateById(ctx, cardId, data)
	if err != nil {
		return nil, err
	}

	if err = uc.cardAgg.Aggregate(ctx, card); err != nil {
		return nil, err
	}

	return card, nil
}
