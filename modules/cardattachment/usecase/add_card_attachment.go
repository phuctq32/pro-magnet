package cauc

import (
	"golang.org/x/net/context"
	camodel "pro-magnet/modules/cardattachment/model"
)

func (uc *cardAttachmentUseCase) AddCardAttachment(
	ctx context.Context,
	data *camodel.CardAttachment,
) (*camodel.CardAttachment, error) {
	// Get board id
	data.BoardId = "654c9df44d947235d14355cc"

	data.Status = camodel.Active

	return uc.caRepo.Create(ctx, data)
}
