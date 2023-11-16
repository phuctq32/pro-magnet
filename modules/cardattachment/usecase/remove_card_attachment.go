package cauc

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"pro-magnet/common"
	cardmodel "pro-magnet/modules/card/model"
	camodel "pro-magnet/modules/cardattachment/model"
)

func (uc *cardAttachmentUseCase) RemoveCardAttachment(
	ctx context.Context,
	cardId, id string,
) error {
	card, err := uc.cardRepo.FindById(ctx, cardId)
	if err != nil {
		return err
	}
	if card.Status == cardmodel.Deleted {
		return common.NewBadRequestErr(cardmodel.ErrCardDeleted)
	}

	cardAttachment, err := uc.caRepo.FindById(ctx, id)
	if err != nil {
		return err
	}
	if cardAttachment.Status == camodel.Deleted {
		return common.NewBadRequestErr(camodel.ErrCardAttachmentDeleted)
	}
	if cardAttachment.CardId != cardId {
		return common.NewBadRequestErr(errors.New("card id not match"))
	}

	return uc.caRepo.DeleteById(ctx, id)
}
