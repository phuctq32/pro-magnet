package carduc

import (
	"context"
	"pro-magnet/components/asyncgroup"
	cardmodel "pro-magnet/modules/card/model"
	camodel "pro-magnet/modules/cardattachment/model"
	labelmodel "pro-magnet/modules/label/model"
)

type CardAttachmentRepo interface {
	ListByCardId(ctx context.Context, status camodel.CardAttachmentStatus, id string) ([]camodel.CardAttachment, error)
}

type cardDataAggregator struct {
	asyng  asyncgroup.AsyncGroup
	caRepo CardAttachmentRepo
}

func NewCardDataAggregator(
	asyng asyncgroup.AsyncGroup,
	caRepo CardAttachmentRepo,
) *cardDataAggregator {
	return &cardDataAggregator{
		asyng:  asyng,
		caRepo: caRepo,
	}
}

func (agg *cardDataAggregator) Aggregate(
	ctx context.Context,
	card *cardmodel.Card,
) error {
	if err := agg.asyng.Process(
		ctx,
		agg.aggregateLabels(card),
		agg.aggregateAttachments(card),
	); err != nil {
		return err
	}

	return nil
}

func (agg *cardDataAggregator) aggregateLabels(
	card *cardmodel.Card,
) func(context.Context) error {
	return func(ctx context.Context) error {
		card.Labels = []labelmodel.Label{}
		return nil
	}
}

func (agg *cardDataAggregator) aggregateAttachments(
	card *cardmodel.Card,
) func(context.Context) error {
	return func(ctx context.Context) error {
		attachments, err := agg.caRepo.ListByCardId(ctx, 1, *card.Id)
		if err != nil {
			return err
		}
		card.Attachments = make([]camodel.CardAttachment, 0)

		card.Attachments = attachments
		return nil
	}
}
