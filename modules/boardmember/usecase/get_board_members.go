package bmuc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	bmmodel "pro-magnet/modules/boardmember/model"
	usermodel "pro-magnet/modules/user/model"
)

func (uc *boardMemberUseCase) GetBoardMembers(
	ctx context.Context,
	requesterId, boardId string,
) ([]usermodel.User, error) {
	_, err := uc.boardRepo.FindById(ctx, boardId)
	if err != nil {
		return nil, err
	}

	isBoardMember, err := uc.bmRepo.IsBoardMember(ctx, boardId, requesterId)
	if err != nil {
		return nil, err
	}
	if !isBoardMember {
		return nil, common.NewBadRequestErr(bmmodel.ErrUserNotBoardMember)
	}

	memberIds, err := uc.bmRepo.FindMemberIdByBoardId(ctx, boardId)

	return uc.userRepo.FindSimpleUsersByIds(ctx, memberIds)
}
