package useruc

import (
	"golang.org/x/net/context"
	usermodel "pro-magnet/modules/user/model"
)

func (uc *userUseCase) GetUser(ctx context.Context, userId string) (*usermodel.User, error) {
	return uc.userRepo.FindById(ctx, userId)
}
