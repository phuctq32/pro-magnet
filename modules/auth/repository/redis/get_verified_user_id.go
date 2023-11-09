package authrepo

import (
	"context"
	"github.com/redis/go-redis/v9"
	"pro-magnet/common"
	authmodel "pro-magnet/modules/auth/model"
)

func (redisRepo *authRedisRepository) GetVerifiedUserId(ctx context.Context, verifiedToken string) (*string, error) {
	userId, err := redisRepo.cli.Get(ctx, authmodel.VerifiedTokenPrefix+verifiedToken).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, common.NewServerErr(err)
	}

	return &userId, nil
}
