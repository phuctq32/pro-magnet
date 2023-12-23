package cardchecklistapi

import (
	"golang.org/x/net/context"
	cardchecklistmodel "pro-magnet/modules/cardchecklist/model"
)

type CardChecklistUseCase interface {
	// Checklist
	CreateChecklist(ctx context.Context, cardId string, data *cardchecklistmodel.CardChecklist) error
	UpdateChecklist(ctx context.Context, cardId, checklistId string, updateData *cardchecklistmodel.CardChecklistUpdate) error
	DeleteChecklist(ctx context.Context, cardId, checklistId string) error

	// Checklist Item
	CreateChecklistItem(ctx context.Context, cardId, checklistId string, data *cardchecklistmodel.ChecklistItem) error
	UpdateChecklistItem(ctx context.Context, cardId, checklistId, itemId string, updateData *cardchecklistmodel.ChecklistItemUpdate) error
	DeleteChecklistItem(ctx context.Context, cardId, checklistId, itemId string) error
}

type cardChecklistHandler struct {
	uc CardChecklistUseCase
}

func NewCardChecklistHandler(uc CardChecklistUseCase) *cardChecklistHandler {
	return &cardChecklistHandler{uc: uc}
}
