package useruc

import (
	"golang.org/x/net/context"
	"pro-magnet/components/hasher"
	usermodel "pro-magnet/modules/user/model"
)

type UserRepository interface {
	FindById(ctx context.Context, id string) (*usermodel.User, error)
	UpdatePasswordById(ctx context.Context, id, password string) error
	UpdateById(ctx context.Context, id string, updateData *usermodel.UserUpdate) (*usermodel.User, error)
}

type userUseCase struct {
	userRepo UserRepository
	hasher   hasher.Hasher
}

func NewUserUseCase(
	userRepo UserRepository,
	hasher hasher.Hasher,
) *userUseCase {
	return &userUseCase{
		userRepo: userRepo,
		hasher:   hasher,
	}
}
