package caapi

import (
	"golang.org/x/net/context"
	camodel "pro-magnet/modules/cardattachment/model"
)

type CardAttachmentUseCase interface {
	AddCardAttachment(ctx context.Context, data *camodel.CardAttachment) (*camodel.CardAttachment, error)
	RemoveCardAttachment(ctx context.Context, cardId, id string) error
}

type cardAttachmentHandler struct {
	uc CardAttachmentUseCase
}

func NewCardAttachmentHandler(uc CardAttachmentUseCase) *cardAttachmentHandler {
	return &cardAttachmentHandler{uc: uc}
}
