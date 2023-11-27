package userapi

import (
	"golang.org/x/net/context"
	usermodel "pro-magnet/modules/user/model"
)

type UserUseCase interface {
	GetUser(ctx context.Context, userId string) (*usermodel.User, error)
	ChangePassword(ctx context.Context, userId string, data *usermodel.UserChangePassword) error
}

type userHandler struct {
	uc UserUseCase
}

func NewUserHandler(uc UserUseCase) *userHandler {
	return &userHandler{uc: uc}
}
