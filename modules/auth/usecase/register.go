package authuc

import (
	"context"
	"pro-magnet/common"
	authmodel "pro-magnet/modules/auth/model"
	usermodel "pro-magnet/modules/user/model"
	"time"
)

func (uc *authUseCase) Register(ctx context.Context, data *authmodel.RegisterUser) error {
	isExisted, err := uc.userRepo.UserExist(ctx, data.Email)
	if err != nil {
		return err
	}
	if isExisted {
		return common.NewBadRequestErr(authmodel.ErrUserExisted)
	}

	hashedPw, err := uc.hasher.Hash(data.Password)
	if err != nil {
		return common.NewServerErr(err)
	}

	now := time.Now()
	newUser := &usermodel.User{
		CreatedAt:   now,
		UpdatedAt:   now,
		Email:       data.Email,
		Name:        data.Name,
		Password:    &hashedPw,
		IsVerified:  false,
		Avatar:      authmodel.DefaultAvatarUrl,
		PhoneNumber: &data.PhoneNumber,
		Birthday:    &data.Birthday,
		Type:        usermodel.InternalUser,
	}

	if err = uc.userRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		userId, e := uc.userRepo.Create(txCtx, newUser)
		if e != nil {
			return e
		}
		newUser.Id = userId

		if e = uc.sendVerificationEmail(txCtx, newUser); e != nil {
			return common.NewServerErr(e)
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
