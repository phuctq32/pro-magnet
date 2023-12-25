package carduc

import (
	"context"
	"pro-magnet/components/asyncgroup"
	cardmodel "pro-magnet/modules/card/model"
	camodel "pro-magnet/modules/cardattachment/model"
	labelmodel "pro-magnet/modules/label/model"
	usermodel "pro-magnet/modules/user/model"
)

type CardAttachmentRepo interface {
	ListByCardId(ctx context.Context, status camodel.CardAttachmentStatus, id string) ([]camodel.CardAttachment, error)
}

type UserRepo interface {
	FindSimpleUserById(ctx context.Context, userId string) (*usermodel.User, error)
	FindSimpleUsersByIds(ctx context.Context, userIds []string) ([]usermodel.User, error)
}

type cardDataAggregator struct {
	asyng    asyncgroup.AsyncGroup
	caRepo   CardAttachmentRepo
	userRepo UserRepo
}

func NewCardDataAggregator(
	asyng asyncgroup.AsyncGroup,
	caRepo CardAttachmentRepo,
	userRepo UserRepo,
) *cardDataAggregator {
	return &cardDataAggregator{
		asyng:    asyng,
		caRepo:   caRepo,
		userRepo: userRepo,
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
		agg.aggregateCommentsAuthor(card),
		agg.aggregateCardMembers(card),
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

func (agg *cardDataAggregator) aggregateCommentsAuthor(
	card *cardmodel.Card,
) func(context.Context) error {
	return func(ctx context.Context) error {
		length := len(card.Comments)
		for i := 0; i < length; i++ {
			author, err := agg.userRepo.FindSimpleUserById(ctx, card.Comments[i].AuthorId)
			if err != nil {
				return err
			}
			card.Comments[i].Author = author
		}

		return nil
	}
}

func (agg *cardDataAggregator) aggregateCardMembers(
	card *cardmodel.Card,
) func(context.Context) error {
	return func(ctx context.Context) error {
		members, err := agg.userRepo.FindSimpleUsersByIds(ctx, card.MemberIds)
		if err != nil {
			return err
		}

		card.Members = members
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
