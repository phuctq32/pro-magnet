package wsmemberuc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	wsmembermodel "pro-magnet/modules/workspacemember/model"
)

func (uc *wsMemberUseCase) AddMembers(
	ctx context.Context, requesterId string,
	data *wsmembermodel.WorkspaceMembersCreate,
) error {
	ws, err := uc.wsRepo.FindById(ctx, data.WorkspaceId)
	if err != nil {
		return err
	}

	if requesterId != ws.OwnerUserId {
		return common.NewBadRequestErr(wsmembermodel.ErrUserNotWorkspaceOwner)
	}

	for _, userId := range data.UserIds {
		isWsMember, err := uc.wsMemberRepo.IsWorkspaceMember(ctx, data.WorkspaceId, userId)
		if err != nil {
			return err
		}

		if isWsMember {
			return common.NewBadRequestErr(wsmembermodel.ErrUserAlreadyWorkspaceMember)
		}
	}

	return uc.wsMemberRepo.CreateMany(ctx, data)
}
