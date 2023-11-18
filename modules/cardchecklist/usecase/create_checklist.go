package cardchecklistuc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	cardmodel "pro-magnet/modules/card/model"
	cardchecklistmodel "pro-magnet/modules/cardchecklist/model"
)

func (uc *cardChecklistUseCase) CreateChecklist(
	ctx context.Context,
	cardId string,
	data *cardchecklistmodel.CardChecklist,
) error {
	card, err := uc.cardRepo.FindById(ctx, cardId)
	if err != nil {
		return err
	}
	if card.Status == cardmodel.Deleted {
		return common.NewBadRequestErr(cardmodel.ErrCardDeleted)
	}

	data.Items = make([]cardchecklistmodel.ChecklistItem, 0)

	return uc.ccRepo.Create(ctx, *card.Id, data)
}
