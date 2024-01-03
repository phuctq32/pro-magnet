package userapi

import (
	"golang.org/x/net/context"
	usermodel "pro-magnet/modules/user/model"
)

type UserUseCase interface {
	GetUser(ctx context.Context, userId string) (*usermodel.User, error)
	ChangePassword(ctx context.Context, userId string, data *usermodel.UserChangePassword) error
	UpdateUser(ctx context.Context, userId string, data *usermodel.UserUpdate) (*usermodel.User, error)

	GetUsersToAddToCard(ctx context.Context, requesterId, cardId string) ([]usermodel.User, error)
	GetUsersToAddToWorkspace(ctx context.Context, requesterId, workspaceId, emailSearchQuery string) ([]usermodel.User, error)

	AddSkils(ctx context.Context, requesterId string, skills []string) error
	RemoveSkill(ctx context.Context, requesterId string, skill string) error
}

type userHandler struct {
	uc UserUseCase
}

func NewUserHandler(uc UserUseCase) *userHandler {
	return &userHandler{uc: uc}
}
