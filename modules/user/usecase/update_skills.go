package useruc

import (
	"golang.org/x/net/context"
)

func (uc *userUseCase) UpdateSkills(
	ctx context.Context,
	requesterId string,
	skills []string,
) error {
	_, err := uc.userRepo.FindById(ctx, requesterId)
	if err != nil {
		return err
	}

	return uc.userRepo.UpdateSkills(ctx, requesterId, skills)
}
