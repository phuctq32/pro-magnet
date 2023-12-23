package cardchecklistuc

import (
	"golang.org/x/net/context"
	cardmodel "pro-magnet/modules/card/model"
	cardchecklistmodel "pro-magnet/modules/cardchecklist/model"
)

type CardRepository interface {
	FindById(ctx context.Context, id string) (*cardmodel.Card, error)
}

type CardChecklistRepository interface {
	// Checklist
	Create(ctx context.Context, cardId string, data *cardchecklistmodel.CardChecklist) error
	Update(ctx context.Context, cardId, checklistId string, data *cardchecklistmodel.CardChecklistUpdate) error
	Delete(ctx context.Context, cardId, checklistId string) error

	// Checklist Item
	CreateChecklistItem(ctx context.Context, cardId, checklistId string, data *cardchecklistmodel.ChecklistItem) error
}

type cardChecklistUseCase struct {
	ccRepo   CardChecklistRepository
	cardRepo CardRepository
}

func NewCardChecklistUseCase(
	ccRepo CardChecklistRepository,
	cardRepo CardRepository,
) *cardChecklistUseCase {
	return &cardChecklistUseCase{
		ccRepo:   ccRepo,
		cardRepo: cardRepo,
	}
}
