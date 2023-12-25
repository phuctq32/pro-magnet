package carduc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	cardmodel "pro-magnet/modules/card/model"
	columnmodel "pro-magnet/modules/column/model"
	"slices"
)

func (uc *cardUseCase) RemoveMemberFromCard(
	ctx context.Context,
	requesterId, cardId, memberId string,
) error {
	card, err := uc.cardRepo.FindById(ctx, cardId)
	if err != nil {
		return err
	}
	if card.Status == cardmodel.Deleted {
		return common.NewBadRequestErr(cardmodel.ErrCardDeleted)
	}

	// Check requester is a member of card's board
	isBoardMember, err := uc.bmRepo.IsBoardMember(ctx, card.BoardId, requesterId)
	if err != nil {
		return err
	}
	if !isBoardMember {
		return common.NewBadRequestErr(columnmodel.ErrNotBoardMember)
	}

	if !slices.Contains(card.MemberIds, memberId) {
		return common.NewBadRequestErr(cardmodel.ErrUserNotExistInCard)
	}

	return uc.cardRepo.RemoveMember(ctx, cardId, memberId)
}
