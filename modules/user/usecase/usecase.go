package useruc

import (
	"golang.org/x/net/context"
	usermodel "pro-magnet/modules/user/model"
)

type UserRepository interface {
	FindById(ctx context.Context, id string) (*usermodel.User, error)
}

type userUseCase struct {
	userRepo UserRepository
}

func NewUserUseCase(
	userRepo UserRepository,
) *userUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}
