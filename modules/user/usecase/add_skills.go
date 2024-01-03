package useruc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	usermodel "pro-magnet/modules/user/model"
)

func (uc *userUseCase) AddSkils(
	ctx context.Context,
	requesterId string,
	skills []string,
) error {
	user, err := uc.userRepo.FindById(ctx, requesterId)
	if err != nil {
		return err
	}

	existedSkillMap := map[string]bool{}
	for _, skill := range user.Skills {
		existedSkillMap[skill] = true
	}

	for _, skill := range skills {
		if existedSkillMap[skill] {
			return common.NewBadRequestErr(usermodel.ErrSkillAlreadyExisted)
		}
	}

	return uc.userRepo.UpdateSkills(ctx, requesterId, skills)
}
