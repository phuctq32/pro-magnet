package wsuc

import (
	"golang.org/x/net/context"
	"pro-magnet/components/asyncgroup"
	usermodel "pro-magnet/modules/user/model"
	wsmodel "pro-magnet/modules/workspace/model"
)

type UserRepo interface {
	FindSimpleUsersByIds(ctx context.Context, userIds []string) ([]usermodel.User, error)
}

type WorkspaceMemberRepo interface {
	FindMemberIdsByWorkspaceId(ctx context.Context, workspaceId string) ([]string, error)
}

type wsAggregator struct {
	asyncg       asyncgroup.AsyncGroup
	userRepo     UserRepo
	wsMemberRepo WorkspaceMemberRepo
}

func NewWorkspaceAggregator(
	asyncg asyncgroup.AsyncGroup,
	userRepo UserRepo,
	wsMemberRepo WorkspaceMemberRepo,
) *wsAggregator {
	return &wsAggregator{
		asyncg:       asyncg,
		userRepo:     userRepo,
		wsMemberRepo: wsMemberRepo,
	}
}

func (wa *wsAggregator) Aggregate(ctx context.Context, ws *wsmodel.Workspace) error {
	return wa.asyncg.Process(
		ctx,
		wa.aggregateMembers(ws),
	)
}

func (wa *wsAggregator) aggregateMembers(ws *wsmodel.Workspace) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		memberIds, err := wa.wsMemberRepo.FindMemberIdsByWorkspaceId(ctx, *ws.Id)
		if err != nil {
			return err
		}

		members, err := wa.userRepo.FindSimpleUsersByIds(ctx, memberIds)
		if err != nil {
			return err
		}
		ws.Members = members

		return nil
	}
}
