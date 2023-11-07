package authrepo

import (
	"context"
	"pro-magnet/common"
	authmodel "pro-magnet/modules/auth/model"
	"time"
)

func (redisRepo *authRedisRepository) SetVerifiedToken(ctx context.Context, token, userId string) error {
	if err := redisRepo.cli.Set(ctx, authmodel.VerifiedTokenPrefix+token, userId, 30*time.Minute).Err(); err != nil {
		return common.NewServerErr(err)
	}

	return nil
}
