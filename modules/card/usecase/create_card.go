package carduc

import (
	"context"
	cardmodel "pro-magnet/modules/card/model"
	camodel "pro-magnet/modules/cardattachment/model"
	labelmodel "pro-magnet/modules/label/model"
)

func (uc *cardUseCase) CreateCard(
	ctx context.Context,
	data *cardmodel.CardCreation,
) (*cardmodel.Card, error) {
	// Check column exist and Get board id
	// Mock
	data.BoardId = "654c9df44d947235d14355cc"

	newCard, err := uc.cardRepo.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	newCard.Labels = []labelmodel.Label{}
	newCard.Attachments = []camodel.CardAttachment{}
	return newCard, nil
}
