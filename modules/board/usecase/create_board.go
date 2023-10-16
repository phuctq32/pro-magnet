package boarduc

import (
	"context"
	boardmodel "pro-magnet/modules/board/model"
	"slices"
)

func (uc *boardUseCase) CreateBoard(
	ctx context.Context,
	data *boardmodel.BoardCreation,
) (*boardmodel.Board, error) {
	checkBoardNameExistedTask := func(ctx context.Context) error {
		ok, err := uc.boardRepo.ExistsInWorkspace(ctx, data.WorkspaceId)
		if err != nil {
			return err
		}
		if ok {
			return boardmodel.ErrExistedBoard
		}
		return nil
	}

	checkBoardAdIsWorkspaceMember := func(ctx context.Context) error {
		wsMemberIds, err := uc.wsRepo.GetMemberIds(ctx, data.WorkspaceId)
		if err != nil {
			return err
		}
		if !slices.Contains(wsMemberIds, data.UserId) {
			return boardmodel.ErrIsNotMemberOfWorkspace
		}
		return nil
	}

	if err := uc.asyncg.Process(ctx,
		checkBoardNameExistedTask,
		checkBoardAdIsWorkspaceMember,
	); err != nil {
		return nil, err
	}

	return uc.boardRepo.Create(ctx, data)
}
