package cardchecklistuc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	cardmodel "pro-magnet/modules/card/model"
	cardchecklistmodel "pro-magnet/modules/cardchecklist/model"
)

func (uc *cardChecklistUseCase) CreateChecklistItem(
	ctx context.Context,
	cardId, checklistId string,
	data *cardchecklistmodel.ChecklistItem,
) error {
	card, err := uc.cardRepo.FindById(ctx, cardId)
	if err != nil {
		return err
	}
	if card.Status == cardmodel.Deleted {
		return common.NewBadRequestErr(cardmodel.ErrCardDeleted)
	}

	isChecklistExist := false
	for i := 0; i < len(card.Checklists); i++ {
		if *card.Checklists[i].Id == checklistId {
			isChecklistExist = true
			break
		}
	}
	if !isChecklistExist {
		return common.NewBadRequestErr(cardchecklistmodel.ErrChecklistNotFound)
	}

	return uc.ccRepo.CreateChecklistItem(ctx, cardId, checklistId, data)
}
