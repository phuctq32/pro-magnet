package useruc

import (
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"pro-magnet/common"
	usermodel "pro-magnet/modules/user/model"
)

func (uc *userUseCase) ChangePassword(
	ctx context.Context,
	userId string,
	data *usermodel.UserChangePassword,
) error {
	user, err := uc.userRepo.FindById(ctx, userId)
	if err != nil {
		return err
	}

	log.Debug().Str("pw", *user.Password).Msg("")

	if !uc.hasher.Compare(*user.Password, data.CurrentPassword) {
		return common.NewBadRequestErr(usermodel.ErrIncorrectPassword)
	}

	hashedPassword, err := uc.hasher.Hash(data.Password)
	if err != nil {
		return err
	}

	return uc.userRepo.UpdatePasswordById(ctx, userId, hashedPassword)
}
