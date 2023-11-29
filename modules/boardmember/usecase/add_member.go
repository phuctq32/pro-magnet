package bmuc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	bmmodel "pro-magnet/modules/boardmember/model"
)

func (uc *boardMemberUseCase) AddMember(
	ctx context.Context,
	requesterId string,
	data *bmmodel.BoardMember,
) error {
	// Check requester is a board member
	isRequesterABoardMember, err := uc.bmRepo.IsBoardMember(ctx, data.BoardId, requesterId)
	if err != nil {
		return err
	}
	if !isRequesterABoardMember {
		return common.NewBadRequestErr(bmmodel.ErrUserNotBoardMember)
	}

	// Check user to add is a board's workspace member, if user has added before return error
	// current none

	// Check user is added to board before, if true return error
	isUserABoardMember, err := uc.bmRepo.IsBoardMember(ctx, data.BoardId, data.UserId)
	if err != nil {
		return err
	}
	if isUserABoardMember {
		return common.NewBadRequestErr(bmmodel.ErrUserIsABoardMemberBefore)
	}

	return uc.bmRepo.Create(ctx, data)
}
