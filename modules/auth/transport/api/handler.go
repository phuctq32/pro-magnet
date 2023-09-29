package authapi

import (
	"context"
	authmodel "pro-magnet/modules/auth/model"
)

type AuthUseCase interface {
	Register(ctx context.Context, data *authmodel.RegisterUser) error
}

type authHandler struct {
	uc AuthUseCase
}

func NewAuthHandler(uc AuthUseCase) *authHandler {
	return &authHandler{uc: uc}
}
