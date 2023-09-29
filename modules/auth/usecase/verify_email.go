package authuc

import (
	"context"
	"pro-magnet/common"
)

func (uc *authUseCase) Verify(ctx context.Context, verifiedToken string) error {
	userId, err := uc.redisRepo.GetVerifiedUserId(ctx, verifiedToken)
	if err != nil {
		return err
	}
	if userId == nil {
		return common.NewBadRequestErr(err, "invalid verified token")
	}

	if err = uc.userRepo.SetEmailVerified(ctx, *userId); err != nil {
		return err
	}

	return nil
}
