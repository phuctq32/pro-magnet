package carduc

import (
	"context"
	cardmodel "pro-magnet/modules/card/model"
	columnmodel "pro-magnet/modules/column/model"
)

type BoardMemberRepository interface {
	IsBoardMember(ctx context.Context, boardId, userId string) (bool, error)
}

type ColumnRepository interface {
	FindById(ctx context.Context, id string) (*columnmodel.Column, error)
	UpdateById(ctx context.Context, id string, updateData *columnmodel.ColumnUpdate) (*columnmodel.Column, error)
}

type CardRepository interface {
	Create(ctx context.Context, data *cardmodel.CardCreation) (*cardmodel.Card, error)
	FindById(ctx context.Context, id string) (*cardmodel.Card, error)
	UpdateById(ctx context.Context, id string, updateData *cardmodel.CardUpdate) (*cardmodel.Card, error)
	UpdateDate(ctx context.Context, id string, updateData *cardmodel.CardDateUpdate) error
	RemoveDate(ctx context.Context, id string) error
	UpdateMembers(ctx context.Context, cardId string, memberId []string) error
	RemoveMember(ctx context.Context, cardId, memberId string) error
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}

type CardDataAggregator interface {
	Aggregate(ctx context.Context, card *cardmodel.Card) error
}

type cardUseCase struct {
	cardRepo CardRepository
	colRepo  ColumnRepository
	bmRepo   BoardMemberRepository
	cardAgg  CardDataAggregator
}

func NewCardUseCase(
	cardRepo CardRepository,
	colRepo ColumnRepository,
	bmRepo BoardMemberRepository,
	cardAgg CardDataAggregator,
) *cardUseCase {
	return &cardUseCase{
		cardRepo: cardRepo,
		colRepo:  colRepo,
		bmRepo:   bmRepo,
		cardAgg:  cardAgg,
	}
}
