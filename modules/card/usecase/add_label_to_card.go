package carduc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	bmmodel "pro-magnet/modules/boardmember/model"
	cardmodel "pro-magnet/modules/card/model"
	labelmodel "pro-magnet/modules/label/model"
)

func (uc *cardUseCase) AddLabelToCard(
	ctx context.Context,
	requesterId, cardId, labelId string,
) error {
	return uc.cardRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		card, err := uc.cardRepo.FindById(ctx, cardId)
		if err != nil {
			return err
		}
		if card.Status == cardmodel.Deleted {
			return common.NewBadRequestErr(cardmodel.ErrCardDeleted)
		}

		isBoardMember, err := uc.bmRepo.IsBoardMember(txCtx, *card.BoardId, requesterId)
		if err != nil {
			return err
		}
		if !isBoardMember {
			return common.NewBadRequestErr(bmmodel.ErrUserNotBoardMember)
		}

		label, err := uc.labelRepo.FindById(txCtx, labelId)
		if err != nil {
			return err
		}
		if label.Status == labelmodel.Deleted {
			return common.NewBadRequestErr(labelmodel.ErrLabelDeleted)
		}

		if label.BoardId != *card.BoardId {
			return common.NewBadRequestErr(labelmodel.ErrLabelNotExistInBoard)
		}

		isLabelExistInCard := false
		for _, id := range card.LabelIds {
			if labelId == id {
				isLabelExistInCard = true
				break
			}
		}
		if isLabelExistInCard {
			return common.NewBadRequestErr(labelmodel.ErrLabelAlreadyExistInCard)
		}

		if err = uc.cardRepo.UpdateLabel(ctx, cardId, labelId); err != nil {
			return err
		}

		return nil
	})
}
