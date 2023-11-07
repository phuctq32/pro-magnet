package authuc

import (
	"context"
	"github.com/pkg/errors"
	"pro-magnet/common"
)

func (uc *authUseCase) ResetPassword(ctx context.Context, resetToken, newPassword string) error {
	email, err := uc.redisRepo.GetResetPasswordEmail(ctx, resetToken)
	if err != nil {
		return err
	}
	if email == nil {
		return common.NewBadRequestErr(errors.New("invalid reset token"))
	}

	hashedPw, err := uc.hasher.Hash(newPassword)
	if err != nil {
		return err
	}

	return uc.userRepo.UpdatePasswordByEmail(ctx, *email, hashedPw)
}
