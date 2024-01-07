package wsuc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	wrkspmodel "pro-magnet/modules/workspace/model"
	wsmembermodel "pro-magnet/modules/workspacemember/model"
)

func (uc *workspaceUseCase) GetWorkspaceById(
	ctx context.Context,
	requesterId, workspaceId string,
) (*wrkspmodel.Workspace, error) {
	ws, err := uc.wsRepo.FindById(ctx, workspaceId)
	if err != nil {
		return nil, err
	}

	isWsMember, err := uc.wsMemberRepo.IsWorkspaceMember(ctx, workspaceId, requesterId)
	if err != nil {
		return nil, err
	}

	if !isWsMember {
		return nil, common.NewBadRequestErr(wsmembermodel.ErrUserNotWorkspaceMember)
	}

	if err = uc.wsAgg.Aggregate(ctx, ws); err != nil {
		return nil, err
	}

	return ws, nil
}
