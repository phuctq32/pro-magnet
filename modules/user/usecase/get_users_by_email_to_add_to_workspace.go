package useruc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	usermodel "pro-magnet/modules/user/model"
	wsmembermodel "pro-magnet/modules/workspacemember/model"
)

func (uc *userUseCase) GetUsersToAddToWorkspace(
	ctx context.Context,
	requesterId, workspaceId, emailSearchQuery string,
) ([]usermodel.User, error) {
	ws, err := uc.wsRepo.FindById(ctx, workspaceId)
	if err != nil {
		return nil, err
	}

	if requesterId != ws.OwnerUserId {
		return nil, common.NewBadRequestErr(wsmembermodel.ErrUserNotWorkspaceOwner)
	}

	wsMemberIds, err := uc.wsMemberRepo.FindMemberIdsByWorkspaceId(ctx, workspaceId)
	if err != nil {
		return nil, err
	}

	return uc.userRepo.SearchUsersByEmail(ctx, emailSearchQuery, wsMemberIds)
}
