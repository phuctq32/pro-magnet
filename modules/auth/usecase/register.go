package authuc

import (
	"context"
	"pro-magnet/common"
	authmodel "pro-magnet/modules/auth/model"
	usermodel "pro-magnet/modules/user/model"
	wsmodel "pro-magnet/modules/workspace/model"
	"time"
)

func (uc *authUseCase) Register(ctx context.Context, data *authmodel.RegisterUser) error {
	isExisted, e := uc.userRepo.UserExist(ctx, data.Email)
	if e != nil {
		return e
	}
	if isExisted {
		return common.NewBadRequestErr(authmodel.ErrUserExisted)
	}

	hashedPw, e := uc.hasher.Hash(data.Password)
	if e != nil {
		return common.NewServerErr(e)
	}

	now := time.Now()
	newUser := &usermodel.User{
		CreatedAt:   &now,
		UpdatedAt:   &now,
		Email:       &data.Email,
		Name:        data.Name,
		Password:    &hashedPw,
		IsVerified:  new(bool),
		Avatar:      authmodel.DefaultAvatarUrl,
		PhoneNumber: &data.PhoneNumber,
		Birthday:    &data.Birthday,
	}
	userType := usermodel.InternalUser
	newUser.Type = &userType

	if e = uc.userRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		userId, e := uc.userRepo.Create(txCtx, newUser)
		if e != nil {
			return e
		}
		newUser.Id = userId

		// Create default workspace for user
		defaultWorkspace := &wsmodel.WorkspaceCreation{
			Name:        newUser.Name + "'s Workspace",
			OwnerUserId: newUser.UserId(),
			Image:       wsmodel.DefaultImageUrl,
		}
		if _, e = uc.wsRepo.Create(ctx, defaultWorkspace); e != nil {
			return e
		}

		if e = uc.sendVerificationEmail(txCtx, newUser); e != nil {
			return common.NewServerErr(e)
		}

		return nil
	}); e != nil {
		return e
	}

	return nil
}
