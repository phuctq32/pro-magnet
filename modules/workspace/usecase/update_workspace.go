package wsuc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	wrkspmodel "pro-magnet/modules/workspace/model"
	wsmembermodel "pro-magnet/modules/workspacemember/model"
)

func (uc *workspaceUseCase) UpdateWorkspace(
	ctx context.Context,
	requesterId, workspaceId string,
	updateData *wrkspmodel.WorkspaceUpdate,
) error {
	return uc.wsRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		ws, err := uc.wsRepo.FindById(ctx, workspaceId)
		if err != nil {
			return err
		}

		isWsMember, err := uc.wsMemberRepo.IsWorkspaceMember(ctx, workspaceId, requesterId)
		if err != nil {
			return err
		}

		if requesterId != ws.OwnerUserId || !isWsMember {
			return common.NewBadRequestErr(wsmembermodel.ErrUserNotWorkspaceMember)
		}

		if err = uc.wsRepo.UpdateById(ctx, workspaceId, updateData); err != nil {
			return err
		}

		return nil
	})
}
