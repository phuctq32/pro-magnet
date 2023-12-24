package cardcommentuc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	cardcommentmodel "pro-magnet/modules/cardcomment/model"
)

func (uc *cardCommentUseCase) DeleteCardComment(
	ctx context.Context,
	requesterId, cardId, commentId string,
) error {
	card, err := uc.validate(ctx, cardId, requesterId)
	if err != nil {
		return err
	}

	isCommentExist := false
	i := 0
	for ; i < len(card.Comments); i++ {
		if *card.Comments[i].Id == commentId {
			isCommentExist = true
			break
		}
	}
	if !isCommentExist {
		return common.NewBadRequestErr(cardcommentmodel.ErrCommentNotFound)
	}
	if card.Comments[i].AuthorId != requesterId {
		return common.NewBadRequestErr(cardcommentmodel.ErrNotCommentAuthor)
	}

	return uc.cmRepo.Delete(ctx, cardId, commentId)
}
