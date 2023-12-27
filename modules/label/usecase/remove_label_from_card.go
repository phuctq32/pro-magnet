package labeluc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	bmmodel "pro-magnet/modules/boardmember/model"
	cardmodel "pro-magnet/modules/card/model"
	labelmodel "pro-magnet/modules/label/model"
)

func (uc *labelUseCase) RemoveLabelFromCard(
	ctx context.Context,
	requesterId, cardId, labelId string,
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

		card, err := uc.cardRepo.FindById(ctx, cardId)
		if err != nil {
			return err
		}
		if card.Status == cardmodel.Deleted {
			return common.NewBadRequestErr(cardmodel.ErrCardDeleted)
		}

		isLabelExistInCard := false
		for _, id := range card.LabelIds {
			if labelId == id {
				isLabelExistInCard = true
				break
			}
		}
		if !isLabelExistInCard {
			return common.NewBadRequestErr(labelmodel.ErrLabelNotExistInCard)
		}

		if err = uc.cardRepo.RemoveLabel(ctx, cardId, labelId); err != nil {
			return err
		}

		return nil
	})
}
