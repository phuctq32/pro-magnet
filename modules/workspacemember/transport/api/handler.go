package wsmemberapi

import (
	"golang.org/x/net/context"
	wsmembermodel "pro-magnet/modules/workspacemember/model"
)

type WorkspaceMemberUseCase interface {
	AddMembers(ctx context.Context, requesterId string, data *wsmembermodel.WorkspaceMembersCreate) error
}

type wsMemberHandler struct {
	uc WorkspaceMemberUseCase
}

func NewWorkspaceMemberHandler(uc WorkspaceMemberUseCase) *wsMemberHandler {
	return &wsMemberHandler{uc: uc}
}
