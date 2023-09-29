package authuc

import (
	"context"
	"github.com/pkg/errors"
	"pro-magnet/common"
)

func (uc *authUseCase) Verify(ctx context.Context, verifiedToken string) error {
	userId, err := uc.redisRepo.GetVerifiedUserId(ctx, verifiedToken)
	if err != nil {
		return err
	}
	if userId == nil {
		return common.NewBadRequestErr(errors.New("invalid token"))
	}

	if err = uc.userRepo.SetEmailVerified(ctx, *userId); err != nil {
		return err
	}

	return nil
}
