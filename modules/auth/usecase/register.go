package authuc

import (
	"context"
	"pro-magnet/common"
	authmodel "pro-magnet/modules/auth/model"
	usermodel "pro-magnet/modules/user/model"
	"time"
)

func (uc *authUseCase) Register(ctx context.Context, data *authmodel.RegisterUser) error {
	// use case should be in a transaction
	if err := uc.userRepo.CheckEmailExists(ctx, data.Email); err != nil {
		return err
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
		Password:    hashedPw,
		IsVerified:  false,
		Avatar:      authmodel.DefaultAvatarUrl,
		PhoneNumber: data.PhoneNumber,
		Birthday:    data.Birthday,
	}

	userId, err := uc.userRepo.Create(ctx, newUser)
	if err != nil {
		return err
	}
	newUser.Id = userId

	if err = uc.sendVerificationEmail(ctx, newUser); err != nil {
		return common.NewServerErr(err)
	}

	return nil
}
