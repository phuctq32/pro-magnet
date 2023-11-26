package columnuc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	columnmodel "pro-magnet/modules/column/model"
)

func (uc *columnUseCase) UpdateColumn(
	ctx context.Context,
	userId, columnId string,
	data *columnmodel.ColumnUpdate,
) (*columnmodel.Column, error) {
	col, err := uc.colRepo.FindById(ctx, columnId)
	if err != nil {
		return nil, err
	}

	isBoardMember, err := uc.bmRepo.IsBoardMember(ctx, col.BoardId, userId)
	if err != nil {
		return nil, err
	}
	if !isBoardMember {
		return nil, common.NewBadRequestErr(columnmodel.ErrNotBoardMember)
	}

	return uc.colRepo.UpdateById(ctx, columnId, data)
}
