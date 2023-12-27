package labeluc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	bmmodel "pro-magnet/modules/boardmember/model"
	labelmodel "pro-magnet/modules/label/model"
)

func (uc *labelUseCase) UpdateLabel(
	ctx context.Context,
	requesterId, labelId string,
	updateData *labelmodel.LabelUpdate,
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

		if err = uc.validateLabel(
			txCtx, label.Id,
			label.BoardId,
			*updateData.Title,
			*updateData.Color,
		); err != nil {
			return err
		}

		if err = uc.labelRepo.UpdateById(txCtx, labelId, updateData); err != nil {
			return err
		}

		return nil
	})
}
