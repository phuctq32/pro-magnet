package boarduc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	boardmodel "pro-magnet/modules/board/model"
	bmmodel "pro-magnet/modules/boardmember/model"
)

func (uc *boardUseCase) GetBoardById(
	ctx context.Context,
	requesterId, boardId string,
	labelIds []string,
) (*boardmodel.Board, error) {
	board, err := uc.boardRepo.FindById(ctx, boardId)
	if err != nil {
		return nil, err
	}
	if board.Status == boardmodel.Deleted {
		return nil, common.NewBadRequestErr(boardmodel.ErrBoardDeleted)
	}
	isBoardMember, e := uc.bmRepo.IsBoardMember(ctx, boardId, requesterId)
	if e != nil {
		return nil, e
	}
	if !isBoardMember {
		return nil, common.NewBadRequestErr(bmmodel.ErrUserNotBoardMember)
	}

	board.FilteredLabelIds = labelIds
	// Check if labels are existed in board

	if err = uc.boardAgg.Aggregate(ctx, board); err != nil {
		return nil, err
	}

	return board, nil
}
