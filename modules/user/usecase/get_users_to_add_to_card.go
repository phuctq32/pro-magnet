package useruc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	cardmodel "pro-magnet/modules/card/model"
	columnmodel "pro-magnet/modules/column/model"
	usermodel "pro-magnet/modules/user/model"
)

func (uc *userUseCase) GetUsersToAddToCard(
	ctx context.Context,
	requesterId, cardId string,
) ([]usermodel.User, error) {
	card, err := uc.cardRepo.FindById(ctx, cardId)
	if err != nil {
		return nil, err
	}
	if card.Status == cardmodel.Deleted {
		return nil, common.NewBadRequestErr(cardmodel.ErrCardDeleted)
	}

	// Check user is a member of card's board
	isBoardMember, err := uc.bmRepo.IsBoardMember(ctx, *card.BoardId, requesterId)
	if err != nil {
		return nil, err
	}
	if !isBoardMember {
		return nil, common.NewBadRequestErr(columnmodel.ErrNotBoardMember)
	}

	boardMemberIds, err := uc.bmRepo.FindMemberIdsByBoardId(ctx, *card.BoardId)
	if err != nil {
		return nil, err
	}

	cardMemberIdsMap := map[string]int{}
	for _, id := range card.MemberIds {
		cardMemberIdsMap[id] = 1
	}

	userIdsToAdd := make([]string, 0)
	for _, id := range boardMemberIds {
		if _, ok := cardMemberIdsMap[id]; !ok {
			userIdsToAdd = append(userIdsToAdd, id)
		}
	}

	return uc.userRepo.FindSimpleUsersByIds(ctx, userIdsToAdd)
}
