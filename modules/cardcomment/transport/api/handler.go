package cardcommentapi

import (
	"golang.org/x/net/context"
	cardcommentmodel "pro-magnet/modules/cardcomment/model"
)

type CardCommentUseCase interface {
	CreateCardComment(ctx context.Context, cardId string, data *cardcommentmodel.CardCommentCreate) error
}

type cardCommentHandler struct {
	uc CardCommentUseCase
}

func NewCardCommentHandler(uc CardCommentUseCase) *cardCommentHandler {
	return &cardCommentHandler{uc: uc}
}
