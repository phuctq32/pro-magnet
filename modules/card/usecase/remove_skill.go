package carduc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	cardmodel "pro-magnet/modules/card/model"
	columnmodel "pro-magnet/modules/column/model"
)

func (uc *cardUseCase) RemoveSkill(
	ctx context.Context,
	requesterId, cardId string,
	skill string,
) error {
	card, err := uc.cardRepo.FindById(ctx, cardId)
	if err != nil {
		return err
	}
	if card.Status == cardmodel.Deleted {
		return common.NewBadRequestErr(cardmodel.ErrCardDeleted)
	}

	isBoardMember, err := uc.bmRepo.IsBoardMember(ctx, *card.BoardId, requesterId)
	if err != nil {
		return err
	}
	if !isBoardMember {
		return common.NewBadRequestErr(columnmodel.ErrNotBoardMember)
	}

	isSkillExisted := false
	for _, sk := range card.Skills {
		if sk == skill {
			isSkillExisted = true
			break
		}
	}
	if !isSkillExisted {
		return common.NewBadRequestErr(cardmodel.ErrSkillNotFound)
	}

	return uc.cardRepo.RemoveSkill(ctx, cardId, skill)
}
