package searchuc

import (
	"golang.org/x/net/context"
	"pro-magnet/components/asyncgroup"
	boardmodel "pro-magnet/modules/board/model"
	cardmodel "pro-magnet/modules/card/model"
	wsmodel "pro-magnet/modules/workspace/model"
)

type WorkspaceRepository interface {
	Search(ctx context.Context, memberId, searchTerm string) ([]wsmodel.Workspace, error)
}

type BoardRepository interface {
	Search(ctx context.Context, memberId, searchTerm string) ([]boardmodel.Board, error)
}

type CardRepository interface {
	Search(ctx context.Context, memberId, searchTerm string) ([]cardmodel.Card, error)
}

type searchUseCase struct {
	wsRepo    WorkspaceRepository
	boardRepo BoardRepository
	careRepo  CardRepository
	asyncg    asyncgroup.AsyncGroup
}

func NewSearchUseCase(
	wsRepo WorkspaceRepository,
	boardRepo BoardRepository,
	cardRepo CardRepository,
	asyncg asyncgroup.AsyncGroup,
) *searchUseCase {
	return &searchUseCase{
		wsRepo:    wsRepo,
		boardRepo: boardRepo,
		careRepo:  cardRepo,
		asyncg:    asyncg,
	}
}
