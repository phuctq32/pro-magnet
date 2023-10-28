package boarduc

import (
	"context"
	"pro-magnet/components/asyncgroup"
	boardmodel "pro-magnet/modules/board/model"
)

type BoardRepository interface {
	Create(ctx context.Context, data *boardmodel.BoardCreation) (*boardmodel.Board, error)
	ExistsInWorkspace(ctx context.Context, boardName, workspaceId string) (bool, error)
}

type WorkspaceRepository interface {
	GetMemberIds(ctx context.Context, workspaceId string) ([]string, error)
}

type boardUseCase struct {
	boardRepo BoardRepository
	wsRepo    WorkspaceRepository
	asyncg    asyncgroup.AsyncGroup
}

func NewBoardUseCase(
	boardRepo BoardRepository,
	wsRepo WorkspaceRepository,
	asyncg asyncgroup.AsyncGroup,
) *boardUseCase {
	return &boardUseCase{
		boardRepo: boardRepo,
		wsRepo:    wsRepo,
		asyncg:    asyncg,
	}
}
