package cardapi

import (
	"context"
	cardmodel "pro-magnet/modules/card/model"
)

type CardUseCase interface {
	CreateCard(ctx context.Context, data *cardmodel.CardCreation) (*cardmodel.Card, error)
}

type cardHandler struct {
	uc CardUseCase
}

func NewCardHandler(uc CardUseCase) *cardHandler {
	return &cardHandler{uc: uc}
}
