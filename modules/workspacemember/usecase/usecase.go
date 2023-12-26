package wsmemberuc

import (
	"golang.org/x/net/context"
	wsmodel "pro-magnet/modules/workspace/model"
	wsmembermodel "pro-magnet/modules/workspacemember/model"
)

type WorkspaceRepository interface {
	FindById(ctx context.Context, id string) (*wsmodel.Workspace, error)
}

type WorkspaceMemberRepository interface {
	IsWorkspaceMember(ctx context.Context, workspaceId, userId string) (bool, error)
	CreateMany(ctx context.Context, data *wsmembermodel.WorkspaceMembersCreate) error
	Delete(ctx context.Context, workspaceId, userId string) error
}

type wsMemberUseCase struct {
	wsMemberRepo WorkspaceMemberRepository
	wsRepo       WorkspaceRepository
}

func NewWorkspaceMemberUseCase(
	wsMemberRepo WorkspaceMemberRepository,
	wsRepo WorkspaceRepository,
) *wsMemberUseCase {
	return &wsMemberUseCase{
		wsMemberRepo: wsMemberRepo,
		wsRepo:       wsRepo,
	}
}
