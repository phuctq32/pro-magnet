package boarduc

import (
	"context"
	boardmodel "pro-magnet/modules/board/model"
	bmmodel "pro-magnet/modules/boardmember/model"
	"slices"
)

func (uc *boardUseCase) CreateBoard(
	ctx context.Context,
	data *boardmodel.BoardCreation,
) (board *boardmodel.Board, err error) {
	checkBoardNameExistedTask := func(ctx context.Context) error {
		ok, err := uc.boardRepo.ExistsInWorkspace(ctx, data.Name, data.WorkspaceId)
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

	if err = uc.asyncg.Process(ctx,
		checkBoardNameExistedTask,
		checkBoardAdIsWorkspaceMember,
	); err != nil {
		return nil, err
	}

	err = uc.boardRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		// Create new board
		newBoard, e := uc.boardRepo.Create(txCtx, data)
		if e != nil {
			return e
		}
		board = newBoard

		// Add admin as a board member
		if e = uc.bmRepo.Create(txCtx, &bmmodel.BoardMember{
			BoardId: *board.Id,
			UserId:  data.UserId,
		}); e != nil {
			return e
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return board, nil
}
