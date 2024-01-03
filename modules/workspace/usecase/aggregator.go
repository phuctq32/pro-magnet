package wsuc

import (
	"golang.org/x/net/context"
	"pro-magnet/components/asyncgroup"
	boardmodel "pro-magnet/modules/board/model"
	usermodel "pro-magnet/modules/user/model"
	wsmodel "pro-magnet/modules/workspace/model"
)

type UserRepo interface {
	FindSimpleUsersByIds(ctx context.Context, userIds []string) ([]usermodel.User, error)
}

type WorkspaceMemberRepo interface {
	FindMemberIdsByWorkspaceId(ctx context.Context, workspaceId string) ([]string, error)
}

type BoardRepo interface {
	FindByWorkspaceId(ctx context.Context, status boardmodel.BoardStatus, workspaceId string) ([]boardmodel.Board, error)
}

type wsAggregator struct {
	asyncg       asyncgroup.AsyncGroup
	userRepo     UserRepo
	wsMemberRepo WorkspaceMemberRepo
	boardRepo    BoardRepo
}

func NewWorkspaceAggregator(
	asyncg asyncgroup.AsyncGroup,
	userRepo UserRepo,
	wsMemberRepo WorkspaceMemberRepo,
	boardRepo BoardRepo,
) *wsAggregator {
	return &wsAggregator{
		asyncg:       asyncg,
		userRepo:     userRepo,
		wsMemberRepo: wsMemberRepo,
		boardRepo:    boardRepo,
	}
}

func (wa *wsAggregator) AggregateMany(ctx context.Context, workspaces []wsmodel.Workspace) error {
	tasks := make([]func(context.Context) error, 0)
	for i := 0; i < len(workspaces); i++ {
		tasks = append(tasks, wa.aggregateMembers(&workspaces[i]))
	}

	return wa.asyncg.Process(ctx, tasks...)
}

func (wa *wsAggregator) Aggregate(ctx context.Context, ws *wsmodel.Workspace) error {
	return wa.asyncg.Process(
		ctx,
		wa.aggregateMembers(ws),
		wa.aggregateBoards(ws),
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

func (wa *wsAggregator) aggregateBoards(ws *wsmodel.Workspace) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		boards, err := wa.boardRepo.FindByWorkspaceId(ctx, boardmodel.Active, *ws.Id)
		if err != nil {
			return err
		}

		ws.Boards = boards

		return nil
	}
}
