package carduc

import (
	"context"
	"pro-magnet/common"
	cardmodel "pro-magnet/modules/card/model"
	camodel "pro-magnet/modules/cardattachment/model"
	columnmodel "pro-magnet/modules/column/model"
	labelmodel "pro-magnet/modules/label/model"
)

func (uc *cardUseCase) CreateCard(
	ctx context.Context,
	userId string,
	data *cardmodel.CardCreation,
) (*cardmodel.Card, error) {
	// Check column exist and Get board id
	col, err := uc.colRepo.FindById(ctx, data.ColumnId)
	if err != nil {
		return nil, err
	}

	isBoardMember, err := uc.bmRepo.IsBoardMember(ctx, col.BoardId, userId)
	if err != nil {
		return nil, err
	}
	if !isBoardMember {
		return nil, common.NewBadRequestErr(columnmodel.ErrNotBoardMember)
	}
	data.BoardId = col.BoardId

	var newCard *cardmodel.Card
	err = uc.cardRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		// Create new card
		newCard, err = uc.cardRepo.Create(ctx, data)
		if err != nil {
			return err
		}
		newCard.Labels = []labelmodel.Label{}
		newCard.Attachments = []camodel.CardAttachment{}

		// Updated orderedCardIds of column
		newOrderedCardIds := append(col.OrderedCardIds, *newCard.Id)
		_, err = uc.colRepo.UpdateById(ctx, *col.Id, &columnmodel.ColumnUpdate{
			Title:          nil,
			OrderedCardIds: newOrderedCardIds,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return newCard, nil
}
