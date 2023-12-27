package columnuc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	boardmodel "pro-magnet/modules/board/model"
	columnmodel "pro-magnet/modules/column/model"
)

func (uc *columnUseCase) CreateColumn(
	ctx context.Context,
	userId string,
	data *columnmodel.ColumnCreate,
) (col *columnmodel.Column, err error) {
	err = uc.colRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		board, e := uc.boardRepo.FindById(ctx, data.BoardId)
		if e != nil {
			return e
		}
		if board.Status == boardmodel.Deleted {
			return common.NewBadRequestErr(boardmodel.ErrBoardDeleted)
		}

		// Check user is a board member
		isBoardMember, e := uc.bmRepo.IsBoardMember(ctx, data.BoardId, userId)
		if e != nil {
			return e
		}
		if !isBoardMember {
			return common.NewBadRequestErr(columnmodel.ErrNotBoardMember)
		}

		data.Status = columnmodel.Active

		col, e = uc.colRepo.Create(ctx, data)
		if e != nil {
			return e
		}

		if e = uc.boardRepo.AddColumnId(ctx, data.BoardId, *col.Id); e != nil {
			return e
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return col, nil
}
