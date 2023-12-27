package boarduc

import (
	"context"
	"pro-magnet/components/asyncgroup"
	boardmodel "pro-magnet/modules/board/model"
	bmmodel "pro-magnet/modules/boardmember/model"
)

type BoardRepository interface {
	Create(ctx context.Context, data *boardmodel.BoardCreation) (*boardmodel.Board, error)
	ExistsInWorkspace(ctx context.Context, boardId *string, boardName, workspaceId string) (bool, error)
	UpdateById(ctx context.Context, boardId string, updateData *boardmodel.BoardUpdate) error
	FindById(ctx context.Context, id string) (*boardmodel.Board, error)
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}

type WorkspaceMemberRepository interface {
	IsWorkspaceMember(ctx context.Context, workspaceId, userId string) (bool, error)
}

type BoardMemberRepository interface {
	Create(ctx context.Context, data *bmmodel.BoardMember) error
	IsBoardMember(ctx context.Context, boardId, userId string) (bool, error)
}

type boardUseCase struct {
	boardRepo    BoardRepository
	bmRepo       BoardMemberRepository
	wsMemberRepo WorkspaceMemberRepository
	asyncg       asyncgroup.AsyncGroup
}

func NewBoardUseCase(
	boardRepo BoardRepository,
	bmRepo BoardMemberRepository,
	wsMemberRepo WorkspaceMemberRepository,
	asyncg asyncgroup.AsyncGroup,
) *boardUseCase {
	return &boardUseCase{
		boardRepo:    boardRepo,
		bmRepo:       bmRepo,
		wsMemberRepo: wsMemberRepo,
		asyncg:       asyncg,
	}
}
