package authuc

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"pro-magnet/common"
	"pro-magnet/components/mailer"
	"pro-magnet/configs"
	usermodel "pro-magnet/modules/user/model"
)

func (uc *authUseCase) SendVerificationEmail(ctx context.Context, email string) error {
	user, err := uc.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return err
	}

	if user.IsVerified == true {
		return common.NewBadRequestErr(errors.New("email already verified"))
	}

	if err = uc.sendVerificationEmail(ctx, user); err != nil {
		return common.NewServerErr(err)
	}

	return nil
}

func (uc *authUseCase) sendVerificationEmail(ctx context.Context, user *usermodel.User) error {
	// generate verified token
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return err
	}
	verifiedToken := hex.EncodeToString(b)

	// set verified token to cache
	if err = uc.redisRepo.SetVerifiedToken(ctx, verifiedToken, *user.Id); err != nil {
		return err
	}

	// send email verification
	emailConfig := mailer.NewEmailConfigWithDynamicTemplate(
		configs.EnvConfigs.SendgridFromEmail(),
		user.Email,
		"Verify Email",
		configs.EnvConfigs.SendgridVerifyEmailTemplateId(),
		map[string]interface{}{
			"username": user.Name,
			"url":      configs.EnvConfigs.VerificationURL() + verifiedToken,
		},
	)

	log.Debug().Msg("sending verification email...")
	if err = uc.mailer.Send(emailConfig); err != nil {
		return err
	}

	return nil
}
