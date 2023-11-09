package authrepo

import (
	"context"
	"pro-magnet/common"
	authmodel "pro-magnet/modules/auth/model"
	"time"
)

func (redisRepo *authRedisRepository) SetResetToken(ctx context.Context, token, email string) error {
	if err := redisRepo.cli.Set(ctx, authmodel.ResetTokenPrefix+token, email, 30*time.Minute).Err(); err != nil {
		return common.NewServerErr(err)
	}

	return nil
}
