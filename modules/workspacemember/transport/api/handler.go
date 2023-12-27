package wsmemberapi

import (
	"golang.org/x/net/context"
	wsmembermodel "pro-magnet/modules/workspacemember/model"
)

type WorkspaceMemberUseCase interface {
	AddMembers(ctx context.Context, requesterId string, data *wsmembermodel.WorkspaceMembersCreate) error
	RemoveMember(ctx context.Context, requesterId, workspaceId, memberId string) error
}

type wsMemberHandler struct {
	uc WorkspaceMemberUseCase
}

func NewWorkspaceMemberHandler(uc WorkspaceMemberUseCase) *wsMemberHandler {
	return &wsMemberHandler{uc: uc}
}
