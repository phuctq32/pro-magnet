package cardchecklistapi

import (
	"golang.org/x/net/context"
	cardchecklistmodel "pro-magnet/modules/cardchecklist/model"
)

type CardChecklistUseCase interface {
	CreateChecklist(ctx context.Context, cardId string, data *cardchecklistmodel.CardChecklist) error
}

type cardChecklistHandler struct {
	uc CardChecklistUseCase
}

func NewCardChecklistHandler(uc CardChecklistUseCase) *cardChecklistHandler {
	return &cardChecklistHandler{uc: uc}
}
