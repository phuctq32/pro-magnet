package carduc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	cardmodel "pro-magnet/modules/card/model"
	columnmodel "pro-magnet/modules/column/model"
)

func (uc *cardUseCase) AddSkils(
	ctx context.Context,
	requesterId, cardId string,
	skills []string,
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

	existedSkillMap := map[string]bool{}
	for _, skill := range card.Skills {
		existedSkillMap[skill] = true
	}

	for _, skill := range skills {
		if existedSkillMap[skill] {
			return common.NewBadRequestErr(cardmodel.ErrSkillAlreadyExisted)
		}
	}

	return uc.cardRepo.UpdateSkills(ctx, cardId, skills)
}
