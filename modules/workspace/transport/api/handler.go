package wsapi

import (
	"context"
	wrkspmodel "pro-magnet/modules/workspace/model"
)

type WorkspaceUseCase interface {
	CreateWorkspace(ctx context.Context, userId string, data *wrkspmodel.WorkspaceCreation) (*wrkspmodel.Workspace, error)
}

type wsHandler struct {
	uc WorkspaceUseCase
}

func NewWorkspaceHandler(uc WorkspaceUseCase) *wsHandler {
	return &wsHandler{uc: uc}
}
