package useruc

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	usermodel "pro-magnet/modules/user/model"
)

func (uc *userUseCase) RemoveSkill(
	ctx context.Context,
	requesterId string,
	skill string,
) error {
	user, err := uc.userRepo.FindById(ctx, requesterId)
	if err != nil {
		return err
	}

	isSkillExisted := false
	for _, sk := range user.Skills {
		if sk == skill {
			isSkillExisted = true
			break
		}
	}
	if !isSkillExisted {
		return common.NewBadRequestErr(usermodel.ErrSkillNotFound)
	}

	return uc.userRepo.RemoveSkill(ctx, requesterId, skill)
}
