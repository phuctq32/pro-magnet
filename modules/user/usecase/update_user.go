package useruc

import (
	"golang.org/x/net/context"
	usermodel "pro-magnet/modules/user/model"
)

func (uc *userUseCase) UpdateUser(
	ctx context.Context,
	userId string,
	data *usermodel.UserUpdate,
) (*usermodel.User, error) {
	return uc.userRepo.UpdateById(ctx, userId, data)
}
