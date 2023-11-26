package carduc

import (
	"context"
	"pro-magnet/common"
	cardmodel "pro-magnet/modules/card/model"
	columnmodel "pro-magnet/modules/column/model"
)

func (uc *cardUseCase) UpdateCardById(
	ctx context.Context,
	userId, cardId string,
	data *cardmodel.CardUpdate,
) (*cardmodel.Card, error) {
	card, err := uc.cardRepo.FindById(ctx, cardId)
	if err != nil {
		return nil, err
	}
	if card.Status == cardmodel.Deleted {
		return nil, common.NewBadRequestErr(cardmodel.ErrCardDeleted)
	}

	// Check user is a member of card's board
	isBoardMember, err := uc.bmRepo.IsBoardMember(ctx, card.BoardId, userId)
	if err != nil {
		return nil, err
	}
	if !isBoardMember {
		return nil, common.NewBadRequestErr(columnmodel.ErrNotBoardMember)
	}

	updatedCard, err := uc.cardRepo.UpdateById(ctx, cardId, data)
	if err != nil {
		return nil, err
	}

	if err = uc.cardAgg.Aggregate(ctx, card); err != nil {
		return nil, err
	}

	return updatedCard, nil
}
