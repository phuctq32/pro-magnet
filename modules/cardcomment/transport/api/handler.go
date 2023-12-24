package cardcommentapi

import (
	"golang.org/x/net/context"
	cardcommentmodel "pro-magnet/modules/cardcomment/model"
)

type CardCommentUseCase interface {
	CreateCardComment(ctx context.Context, cardId string, data *cardcommentmodel.CardCommentCreate) error
	UpdateCardComment(ctx context.Context, requesterId, cardId, commentId string, updateData *cardcommentmodel.CardCommentUpdate) error
	DeleteCardComment(ctx context.Context, requesterId, cardId, commentId string) error
}

type cardCommentHandler struct {
	uc CardCommentUseCase
}

func NewCardCommentHandler(uc CardCommentUseCase) *cardCommentHandler {
	return &cardCommentHandler{uc: uc}
}
