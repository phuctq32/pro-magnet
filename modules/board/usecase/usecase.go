package boarduc

import (
	"context"
	boardmodel "pro-magnet/modules/board/model"
)

type BoardRepository interface {
	Create(ctx context.Context, data *boardmodel.BoardCreation) (*boardmodel.Board, error)
	ExistsInWorkspace(ctx context.Context, workspaceId string) (bool, error)
}

type WorkspaceRepository interface {
	GetMemberIds(ctx context.Context, workspaceId string) ([]string, error)
}

type boardUseCase struct {
	boardRepo BoardRepository
	wsRepo    WorkspaceRepository
}

func NewBoardUseCase(
	boardRepo BoardRepository,
	wsRepo WorkspaceRepository,
) *boardUseCase {
	return &boardUseCase{
		boardRepo: boardRepo,
		wsRepo:    wsRepo,
	}
}
