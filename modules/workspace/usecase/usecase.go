package wsuc

import (
	"context"
	wsmodel "pro-magnet/modules/workspace/model"
	wsmembermodel "pro-magnet/modules/workspacemember/model"
)

type WorkspaceRepository interface {
	Create(context.Context, *wsmodel.WorkspaceCreation) (*wsmodel.Workspace, error)
	FindByName(ctx context.Context, name string) (*wsmodel.Workspace, error)
	FindById(ctx context.Context, workspaceId string) (*wsmodel.Workspace, error)
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}

type WorkspaceMemberRepository interface {
	CreateMany(ctx context.Context, data *wsmembermodel.WorkspaceMembersCreate) error
	IsWorkspaceMember(ctx context.Context, workspaceId, userId string) (bool, error)
}

type WorkspaceAggregator interface {
	Aggregate(ctx context.Context, ws *wsmodel.Workspace) error
}

type workspaceUseCase struct {
	wsRepo       WorkspaceRepository
	wsMemberRepo WorkspaceMemberRepository
	wsAgg        WorkspaceAggregator
}

func NewWorkspaceUseCase(
	wsRepo WorkspaceRepository,
	wsMemberRepo WorkspaceMemberRepository,
	wsAgg WorkspaceAggregator,
) *workspaceUseCase {
	return &workspaceUseCase{
		wsRepo:       wsRepo,
		wsMemberRepo: wsMemberRepo,
		wsAgg:        wsAgg,
	}
}
