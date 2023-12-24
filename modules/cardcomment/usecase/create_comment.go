package cardcommentuc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	cardmodel "pro-magnet/modules/card/model"
	cardcommentmodel "pro-magnet/modules/cardcomment/model"
	columnmodel "pro-magnet/modules/column/model"
)

func (uc *cardCommentUseCase) CreateCardComment(
	ctx context.Context,
	cardId string,
	data *cardcommentmodel.CardCommentCreate,
) error {
	if _, err := uc.validate(ctx, cardId, data.UserId); err != nil {
		return err
	}

	return uc.cmRepo.Create(ctx, cardId, data)
}

func (uc *cardCommentUseCase) validate(
	ctx context.Context,
	cardId, userId string,
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

	return card, nil
}
