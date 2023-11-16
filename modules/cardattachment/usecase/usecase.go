package cauc

import (
	"context"
	cardmodel "pro-magnet/modules/card/model"
	camodel "pro-magnet/modules/cardattachment/model"
)

type CardRepository interface {
	FindById(ctx context.Context, cardId string) (*cardmodel.Card, error)
}

type CardAttachmentRepository interface {
	Create(ctx context.Context, data *camodel.CardAttachment) (*camodel.CardAttachment, error)
	DeleteById(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (*camodel.CardAttachment, error)
}

type cardAttachmentUseCase struct {
	caRepo   CardAttachmentRepository
	cardRepo CardRepository
}

func NewCardAttachmentUseCase(
	caRepo CardAttachmentRepository,
	cardRepo CardRepository,
) *cardAttachmentUseCase {
	return &cardAttachmentUseCase{
		caRepo:   caRepo,
		cardRepo: cardRepo,
	}
}
