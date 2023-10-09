package wsuc

import (
	"context"
	wsmodel "pro-magnet/modules/workspace/model"
)

type WorkspaceRepository interface {
	Create(context.Context, *wsmodel.WorkspaceCreation) (*wsmodel.Workspace, error)
	FindByName(ctx context.Context, name string) (*wsmodel.Workspace, error)
}

type workspaceUseCase struct {
	wsRepo WorkspaceRepository
}

func NewWorkspaceUseCase(wsRepo WorkspaceRepository) *workspaceUseCase {
	return &workspaceUseCase{wsRepo: wsRepo}
}
