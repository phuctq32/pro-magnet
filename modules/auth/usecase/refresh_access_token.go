package authuc

import (
	"context"
	"pro-magnet/common"
)

func (uc *authUseCase) RefreshAccessToken(ctx context.Context, refreshToken string) (*string, error) {
	payload, err := uc.tokenProvider.Validate(refreshToken, uc.rtSecret)
	if err != nil {
		return nil, common.NewUnauthorizedErr(err, "invalid token")
	}

	// Check if user is existing
	if _, err = uc.userRepo.FindById(ctx, payload.UserId); err != nil {
		return nil, err
	}

	accessToken, err := uc.tokenProvider.Generate(payload, uc.atSecret, uc.atExpiry)
	if err != nil {
		return nil, common.NewServerErr(err)
	}

	return accessToken, nil
}
