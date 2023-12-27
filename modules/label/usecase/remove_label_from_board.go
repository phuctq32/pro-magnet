package labeluc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	bmmodel "pro-magnet/modules/boardmember/model"
	labelmodel "pro-magnet/modules/label/model"
)

func (uc *labelUseCase) RemoveLabelFromBoard(
	ctx context.Context,
	requesterId, labelId string,
) error {
	return uc.labelRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		label, err := uc.labelRepo.FindById(txCtx, labelId)
		if err != nil {
			return err
		}
		if label.Status == labelmodel.Deleted {
			return common.NewBadRequestErr(labelmodel.ErrLabelDeleted)
		}

		isBoardMember, err := uc.bmRepo.IsBoardMember(txCtx, label.BoardId, requesterId)
		if err != nil {
			return err
		}
		if !isBoardMember {
			return common.NewBadRequestErr(bmmodel.ErrUserNotBoardMember)
		}

		cardIds, err := uc.cardRepo.FindCardIdsByLabelId(ctx, *label.Id)
		if err != nil {
			return err
		}

		if err = uc.cardRepo.RemoveLabelFromCardsByIds(ctx, cardIds, labelId); err != nil {
			return err
		}

		if err = uc.labelRepo.DeleteById(ctx, labelId); err != nil {
			return err
		}

		return nil
	})
}
