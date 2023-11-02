package authuc

import (
	"context"
	"github.com/pkg/errors"
	"pro-magnet/common"
	"pro-magnet/components/jwt"
	authmodel "pro-magnet/modules/auth/model"
	usermodel "pro-magnet/modules/user/model"
)

func (uc *authUseCase) Login(ctx context.Context, data *authmodel.LoginUser) (*authmodel.LoginResponse, error) {
	user, err := uc.userRepo.FindByEmail(ctx, data.Email)
	if err != nil {
		return nil, err
	}

	if !uc.hasher.Compare(*user.Password, data.Password) {
		return nil, common.NewBadRequestErr(errors.New("email or password invalid"))
	}

	if !user.IsVerified {
		return nil, common.NewBadRequestErr(errors.New("user not verified"))
	}

	accessToken, refreshToken, err := uc.generateTokenPair(ctx, user)
	if err != nil {
		return nil, err
	}

	return &authmodel.LoginResponse{
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
		User:         user,
	}, nil
}

func (uc *authUseCase) generateTokenPair(ctx context.Context, user *usermodel.User) (*string, *string, error) {
	tokenPayload := &jwt.Payload{UserId: *user.Id}

	accessToken, err := uc.tokenProvider.Generate(tokenPayload, uc.atSecret, uc.atExpiry)
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := uc.tokenProvider.Generate(tokenPayload, uc.rtSecret, uc.rtExpiry)
	if err != nil {
		return nil, nil, err
	}

	return accessToken, refreshToken, nil
}
