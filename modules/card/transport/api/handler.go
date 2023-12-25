package cardapi

import (
	"context"
	cardmodel "pro-magnet/modules/card/model"
)

type CardUseCase interface {
	CreateCard(ctx context.Context, userId string, data *cardmodel.CardCreation) (*cardmodel.Card, error)
	GetCardById(ctx context.Context, id string) (*cardmodel.Card, error)
	UpdateCardById(ctx context.Context, userId, cardId string, data *cardmodel.CardUpdate) (*cardmodel.Card, error)
	UpdateCardDate(ctx context.Context, requesterId, cardId string, data *cardmodel.CardDateUpdate) error
	RemoveCardDate(ctx context.Context, requesterId, cardId string) error
	AddMemberToCard(ctx context.Context, requesterId, cardId, memberId string) error
}

type cardHandler struct {
	uc CardUseCase
}

func NewCardHandler(uc CardUseCase) *cardHandler {
	return &cardHandler{uc: uc}
}
