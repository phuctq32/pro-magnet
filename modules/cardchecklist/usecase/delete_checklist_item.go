package cardchecklistuc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	cardmodel "pro-magnet/modules/card/model"
	cardchecklistmodel "pro-magnet/modules/cardchecklist/model"
)

func (uc *cardChecklistUseCase) DeleteChecklistItem(
	ctx context.Context,
	cardId, checklistId, itemId string,
) error {
	card, err := uc.cardRepo.FindById(ctx, cardId)
	if err != nil {
		return err
	}
	if card.Status == cardmodel.Deleted {
		return common.NewBadRequestErr(cardmodel.ErrCardDeleted)
	}

	isChecklistExist := false
	checklistIndex := 0
	for i := 0; i < len(card.Checklists); i++ {
		if *card.Checklists[i].Id == checklistId {
			isChecklistExist = true
			checklistIndex = i
			break
		}
	}
	if !isChecklistExist {
		return common.NewBadRequestErr(cardchecklistmodel.ErrChecklistNotFound)
	}

	isItemExist := false
	for i := 0; i < len(card.Checklists[checklistIndex].Items); i++ {
		if *card.Checklists[checklistIndex].Items[i].Id == itemId {
			isItemExist = true
			break
		}
	}
	if !isItemExist {
		return common.NewBadRequestErr(cardchecklistmodel.ErrChecklistItemNotFound)
	}

	return uc.ccRepo.DeleteChecklistItem(ctx, cardId, checklistId, itemId)
}
