package authuc

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"github.com/rs/zerolog/log"
	"pro-magnet/common"
	"pro-magnet/components/mailer"
	"pro-magnet/configs"
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

	// generate verified token
	b := make([]byte, 32)
	_, err = rand.Read(b)
	if err != nil {
		return common.NewServerErr(err)
	}
	verifiedToken := hex.EncodeToString(b)

	// set verified token to cache
	if err = uc.authRedisRepo.SetVerifiedToken(ctx, verifiedToken, *userId); err != nil {
		return err
	}

	// send email verification
	emailConfig := mailer.NewEmailConfigWithDynamicTemplate(
		configs.EnvConfigs.SendgridFromEmail(),
		newUser.Email,
		"Verify Email",
		configs.EnvConfigs.SendgridVerifyEmailTemplateId(),
		map[string]interface{}{
			"username": newUser.Name,
			"url":      configs.EnvConfigs.VerificationLink() + verifiedToken,
		},
	)
	log.Debug().Msg("sending verification email...")
	if err = uc.mailer.Send(emailConfig); err != nil {
		return common.NewServerErr(err)
	}
	return nil
}
