package authuc

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"github.com/rs/zerolog/log"
	"pro-magnet/common"
	"pro-magnet/components/mailer"
	usermodel "pro-magnet/modules/user/model"
)

func (uc *authUseCase) ForgotPassword(ctx context.Context, email string) error {
	user, err := uc.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return err
	}

	if *user.Type != usermodel.InternalUser {
		return common.NewNotFoundErr("user", usermodel.ErrUserNotFound)
	}

	if !*user.IsVerified {
		return common.NewBadRequestErr(usermodel.ErrUserNotVerified)
	}

	b := make([]byte, 32)
	if _, err = rand.Read(b); err != nil {
		return err
	}
	resetToken := hex.EncodeToString(b)

	log.Debug().Str("resetToken", resetToken).Msg("")

	if err = uc.redisRepo.SetResetToken(ctx, resetToken, email); err != nil {
		return err
	}

	// Send reset password email
	emailConfig := mailer.NewEmailConfigWithDynamicTemplate(
		uc.fromEmail,
		*user.Email,
		"Reset password",
		uc.resetPasswordEmailTemplateId,
		map[string]interface{}{
			"url": uc.resetPasswordUrl + resetToken,
		},
	)

	log.Debug().Msg("sending reset password email...")
	if err = uc.mailer.Send(emailConfig); err != nil {
		return err
	}

	return nil
}
