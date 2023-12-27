package boarduc

import (
	"context"
	boardmodel "pro-magnet/modules/board/model"
	bmmodel "pro-magnet/modules/boardmember/model"
)

func (uc *boardUseCase) CreateBoard(
	ctx context.Context,
	data *boardmodel.BoardCreation,
) (board *boardmodel.Board, err error) {
	checkBoardNameExistedTask := func(ctx context.Context) error {
		ok, e := uc.boardRepo.ExistsInWorkspace(ctx, nil, data.Name, data.WorkspaceId)
		if e != nil {
			return e
		}
		if ok {
			return boardmodel.ErrExistedBoard
		}
		return nil
	}

	checkBoardAdIsWorkspaceMember := func(ctx context.Context) error {
		isWsMember, e := uc.wsMemberRepo.IsWorkspaceMember(ctx, data.WorkspaceId, data.UserId)
		if e != nil {
			return e
		}
		if !isWsMember {
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
