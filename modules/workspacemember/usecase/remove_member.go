package wsmemberuc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	wsmembermodel "pro-magnet/modules/workspacemember/model"
)

func (uc *wsMemberUseCase) RemoveMember(
	ctx context.Context,
	requesterId, workspaceId, memberId string,
) error {
	ws, err := uc.wsRepo.FindById(ctx, workspaceId)
	if err != nil {
		return err
	}

	if requesterId != ws.OwnerUserId {
		return common.NewBadRequestErr(wsmembermodel.ErrUserNotWorkspaceOwner)
	}

	if memberId == ws.OwnerUserId {
		return common.NewBadRequestErr(wsmembermodel.ErrCanNotRemoveWorkspaceOwner)
	}

	isWsMember, err := uc.wsMemberRepo.IsWorkspaceMember(ctx, workspaceId, memberId)
	if err != nil {
		return err
	}
	if !isWsMember {
		return common.NewBadRequestErr(wsmembermodel.ErrUserNotWorkspaceMember)
	}

	return uc.wsMemberRepo.Delete(ctx, workspaceId, memberId)
}
