package carduc

import (
	"context"
	"pro-magnet/components/asyncgroup"
	cardmodel "pro-magnet/modules/card/model"
	labelmodel "pro-magnet/modules/label/model"
)

type CardDataAggregator interface {
	Aggregate(ctx context.Context, card *cardmodel.Card) error
}

type cardDataAggregator struct {
	asyng asyncgroup.AsyncGroup
}

func NewCardDataAggregator(asyng asyncgroup.AsyncGroup) *cardDataAggregator {
	return &cardDataAggregator{
		asyng: asyng,
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
		card.Attachments = []cardmodel.CardAttachment{}
		return nil
	}
}
