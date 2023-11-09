package authrepo

import (
	"context"
	"github.com/redis/go-redis/v9"
	"pro-magnet/common"
	authmodel "pro-magnet/modules/auth/model"
)

func (redisRepo *authRedisRepository) GetResetPasswordEmail(ctx context.Context, resetToken string) (*string, error) {
	email, err := redisRepo.cli.Get(ctx, authmodel.ResetTokenPrefix+resetToken).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, common.NewServerErr(err)
	}

	return &email, nil
}
