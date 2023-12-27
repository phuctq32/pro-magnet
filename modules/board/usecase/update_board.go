package boarduc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	boardmodel "pro-magnet/modules/board/model"
	bmmodel "pro-magnet/modules/boardmember/model"
)

func (uc *boardUseCase) UpdateBoard(
	ctx context.Context,
	requesterId, boardId string,
	updateData *boardmodel.BoardUpdate,
) error {
	return uc.boardRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		board, err := uc.boardRepo.FindById(ctx, boardId)
		if err != nil {
			return err
		}
		if board.Status == boardmodel.Deleted {
			return common.NewBadRequestErr(boardmodel.ErrBoardDeleted)
		}

		// Check requester is board admin
		if updateData.Name != nil {
			if requesterId != board.AdminId {
				return common.NewBadRequestErr(boardmodel.ErrUserNotBoardAdmin)
			}
			isBoardNameExistInWs, err := uc.boardRepo.ExistsInWorkspace(
				ctx, board.Id, *updateData.Name, board.WorkspaceId,
			)
			if err != nil {
				return err
			}
			if isBoardNameExistInWs {
				return boardmodel.ErrExistedBoard
			}
		}

		isBoardMember, err := uc.bmRepo.IsBoardMember(ctx, boardId, requesterId)
		if err != nil {
			return err
		}
		if !isBoardMember {
			return common.NewBadRequestErr(bmmodel.ErrUserNotBoardMember)
		}

		return uc.boardRepo.UpdateById(ctx, boardId, updateData)
	})
}
