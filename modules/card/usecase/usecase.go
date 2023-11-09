package carduc

import (
	"context"
	cardmodel "pro-magnet/modules/card/model"
)

type CardRepository interface {
	Create(ctx context.Context, data *cardmodel.CardCreation) (*cardmodel.Card, error)
}

type cardUseCase struct {
	cardRepo CardRepository
}

func NewCardUseCase(
	cardRepo CardRepository,
) *cardUseCase {
	return &cardUseCase{
		cardRepo: cardRepo,
	}
}
