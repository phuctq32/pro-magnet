package cauc

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"pro-magnet/common"
	camodel "pro-magnet/modules/cardattachment/model"
)

func (uc *cardAttachmentUseCase) RemoveCardAttachment(
	ctx context.Context,
	userId, cardId, cardAttachmentId string,
) error {
	if err := uc.validate(ctx, userId, cardId); err != nil {
		return err
	}

	cardAttachment, err := uc.caRepo.FindById(ctx, cardAttachmentId)
	if err != nil {
		return err
	}
	if cardAttachment.Status == camodel.Deleted {
		return common.NewBadRequestErr(camodel.ErrCardAttachmentDeleted)
	}
	if cardAttachment.CardId != cardId {
		return common.NewBadRequestErr(errors.New("card id not match"))
	}

	return uc.caRepo.DeleteById(ctx, cardAttachmentId)
}
