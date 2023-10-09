package wrkspusecase

import (
	"context"
	"pro-magnet/common"
	wrkspmodel "pro-magnet/modules/workspace/model"
	"time"
)

func (uc *workspaceUseCase) CreateWorkspace(
	ctx context.Context,
	userId string,
	data *wrkspmodel.WorkspaceCreation,
) (*wrkspmodel.Workspace, error) {
	existedWs, err := uc.wsRepo.FindByName(ctx, data.Name)
	if err != nil && err.Error() != common.ErrNotFound.Error() {
		return nil, err
	}
	if existedWs != nil {
		return nil, common.NewBadRequestErr(wrkspmodel.ErrExistedWorkspace)
	}

	now := time.Now()
	newWs := &wrkspmodel.Workspace{
		CreatedAt: now,
		UpdatedAt: now,
		Name:      data.Name,
		UserId:    userId,
		Image:     wrkspmodel.DefaultImageUrl,
	}

	return uc.wsRepo.Create(ctx, newWs)
}
