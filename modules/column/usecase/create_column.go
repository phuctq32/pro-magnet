package columnuc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	columnmodel "pro-magnet/modules/column/model"
)

func (uc *columnUseCase) CreateColumn(
	ctx context.Context,
	userId string,
	data *columnmodel.ColumnCreate,
) (*columnmodel.Column, error) {
	// Check user is a board member
	isBoardMember, err := uc.bmRepo.IsBoardMember(ctx, data.BoardId, userId)
	if err != nil {
		return nil, err
	}
	if !isBoardMember {
		return nil, common.NewBadRequestErr(columnmodel.ErrNotBoardMember)
	}

	data.Status = columnmodel.Active

	return uc.colRepo.Create(ctx, data)
}
