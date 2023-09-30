package authuc

import (
	"context"
	"github.com/pkg/errors"
	"pro-magnet/common"
	"pro-magnet/components/jwt"
	"pro-magnet/configs"
	authmodel "pro-magnet/modules/auth/model"
)

func (uc *authUseCase) Login(ctx context.Context, data *authmodel.LoginUser) (*authmodel.TokenPair, error) {
	user, err := uc.userRepo.FindByEmail(ctx, data.Email)
	if err != nil {
		return nil, err
	}

	if !uc.hasher.Compare(user.Password, data.Password) {
		return nil, common.NewBadRequestErr(errors.New("email or password invalid"))
	}

	if !user.IsVerified {
		return nil, common.NewBadRequestErr(errors.New("user not verified"))
	}

	tokenPayload := &jwt.Payload{UserId: *user.Id}
	accessToken, err := uc.tokenProvider.Generate(
		tokenPayload,
		configs.EnvConfigs.AccessSecret(),
		uc.atExpiry,
	)
	if err != nil {
		return nil, err
	}

	refreshToken, err := uc.tokenProvider.Generate(
		tokenPayload,
		configs.EnvConfigs.RefreshSecret(),
		uc.rtExpiry,
	)

	if err != nil {
		return nil, err
	}

	return &authmodel.TokenPair{
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
	}, nil
}
