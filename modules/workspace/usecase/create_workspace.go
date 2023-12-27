package wsuc

import (
	"context"
	"pro-magnet/common"
	wrkspmodel "pro-magnet/modules/workspace/model"
	wsmembermodel "pro-magnet/modules/workspacemember/model"
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

	data.OwnerUserId = userId
	data.Image = wrkspmodel.DefaultImageUrl

	var newWs *wrkspmodel.Workspace
	err = uc.wsRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		var e error
		newWs, e = uc.wsRepo.Create(ctx, data)
		if e != nil {
			return e
		}

		// Add user to workspace member
		wsMember := &wsmembermodel.WorkspaceMembersCreate{
			WorkspaceId: *newWs.Id,
			UserIds:     []string{data.OwnerUserId},
		}
		if e = uc.wsMemberRepo.CreateMany(ctx, wsMember); e != nil {
			return e
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return newWs, nil
}
