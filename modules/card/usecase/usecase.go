package carduc

import (
	"context"
	cardmodel "pro-magnet/modules/card/model"
)

type CardRepository interface {
	Create(ctx context.Context, data *cardmodel.CardCreation) (*cardmodel.Card, error)
	FindById(ctx context.Context, id string) (*cardmodel.Card, error)
}

type cardUseCase struct {
	cardRepo CardRepository
	cardAgg  CardDataAggregator
}

func NewCardUseCase(
	cardRepo CardRepository,
	cardAgg CardDataAggregator,
) *cardUseCase {
	return &cardUseCase{
		cardRepo: cardRepo,
		cardAgg:  cardAgg,
	}
}
