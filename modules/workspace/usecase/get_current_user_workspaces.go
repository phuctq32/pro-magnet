package wsuc

import (
	"golang.org/x/net/context"
	wrkspmodel "pro-magnet/modules/workspace/model"
)

func (uc *workspaceUseCase) GetCurrentUserWorkspaces(
	ctx context.Context, requesterId string,
) ([]wrkspmodel.Workspace, error) {
	workspaceIds, err := uc.wsMemberRepo.FindWorkspaceIdsByMemberId(ctx, requesterId)
	if err != nil {
		return nil, err
	}

	workspaces, err := uc.wsRepo.FindByIds(ctx, workspaceIds)

	if err = uc.wsAgg.AggregateMany(ctx, workspaces); err != nil {
		return nil, err
	}

	return workspaces, err
}
