package bmuc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	bmmodel "pro-magnet/modules/boardmember/model"
)

func (uc *boardMemberUseCase) AddMember(
	ctx context.Context,
	requesterId string,
	data *bmmodel.AddBoardMembers,
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
	var checkBoardMemberTasks []func(context.Context) error
	for i := 0; i < len(data.UserIds); i++ {
		userId := data.UserIds[i]
		checkBoardMemberTasks = append(
			checkBoardMemberTasks,
			func(ctx context.Context) error {
				isUserABoardMember, err := uc.bmRepo.IsBoardMember(ctx, data.BoardId, userId)
				if err != nil {
					return err
				}
				if isUserABoardMember {
					return common.NewBadRequestErr(bmmodel.ErrUserIsABoardMemberBefore)
				}
				return nil
			})
	}

	if err = uc.asyncg.Process(ctx, checkBoardMemberTasks...); err != nil {
		return err
	}

	return uc.bmRepo.CreateMany(ctx, data)
}
