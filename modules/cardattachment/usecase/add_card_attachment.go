package cauc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	cardmodel "pro-magnet/modules/card/model"
	camodel "pro-magnet/modules/cardattachment/model"
	columnmodel "pro-magnet/modules/column/model"
)

func (uc *cardAttachmentUseCase) AddCardAttachment(
	ctx context.Context,
	userId string,
	data *camodel.CardAttachment,
) (*camodel.CardAttachment, error) {
	if err := uc.validate(ctx, userId, data.CardId); err != nil {
		return nil, err
	}

	data.Status = camodel.Active

	return uc.caRepo.Create(ctx, data)
}

func (uc *cardAttachmentUseCase) validate(
	ctx context.Context,
	userId, cardId string,
) error {
	card, err := uc.cardRepo.FindById(ctx, cardId)
	if err != nil {
		return err
	}
	if card.Status == camodel.Deleted {
		return common.NewNotFoundErr("card", cardmodel.ErrCardDeleted)
	}

	// Check user is a member of card's board
	isBoardMember, err := uc.bmRepo.IsBoardMember(ctx, card.BoardId, userId)
	if err != nil {
		return err
	}
	if !isBoardMember {
		return common.NewBadRequestErr(columnmodel.ErrNotBoardMember)
	}

	return nil
}
