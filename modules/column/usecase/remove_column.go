package columnuc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	columnmodel "pro-magnet/modules/column/model"
)

func (uc *columnUseCase) RemoveColumn(ctx context.Context, userId string, columnId string) error {
	col, err := uc.colRepo.FindById(ctx, columnId)
	if err != nil {
		return err
	}

	isBoardMember, err := uc.bmRepo.IsBoardMember(ctx, col.BoardId, userId)
	if err != nil {
		return err
	}
	if !isBoardMember {
		return common.NewBadRequestErr(columnmodel.ErrNotBoardMember)
	}

	return uc.colRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		if err := uc.cardRepo.DeleteByIds(txCtx, col.OrderedCardIds); err != nil {
			return err
		}

		if err := uc.colRepo.DeleteById(txCtx, columnId); err != nil {
			return err
		}

		return nil
	})
}
