package carduc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	cardmodel "pro-magnet/modules/card/model"
	columnmodel "pro-magnet/modules/column/model"
	"slices"
)

func (uc *cardUseCase) AddMemberToCard(
	ctx context.Context,
	requesterId, cardId string,
	memberIds []string,
) error {
	card, err := uc.cardRepo.FindById(ctx, cardId)
	if err != nil {
		return err
	}
	if card.Status == cardmodel.Deleted {
		return common.NewBadRequestErr(cardmodel.ErrCardDeleted)
	}

	// Check requester and user are members of card's board
	isBoardMember, err := uc.bmRepo.IsBoardMember(ctx, card.BoardId, requesterId)
	if err != nil {
		return err
	}
	if !isBoardMember {
		return common.NewBadRequestErr(columnmodel.ErrNotBoardMember)
	}

	for _, id := range memberIds {
		isBoardMember, err = uc.bmRepo.IsBoardMember(ctx, card.BoardId, id)
		if err != nil {
			return err
		}
		if !isBoardMember {
			return common.NewBadRequestErr(columnmodel.ErrNotBoardMember)
		}

		if slices.Contains(card.MemberIds, id) {
			return common.NewBadRequestErr(cardmodel.ErrUserAddedToCardBefore)
		}
	}

	return uc.cardRepo.UpdateMembers(ctx, cardId, memberIds)
}
