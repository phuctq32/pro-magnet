package authrepo

import (
	"context"
	"pro-magnet/common"
	"time"
)

const VerifiedTokenPrefix string = "verified-token:"

func (redisRepo *authRedisRepository) SetVerifiedToken(ctx context.Context, token, userId string) error {
	if err := redisRepo.cli.Set(ctx, VerifiedTokenPrefix+token, userId, 30*time.Minute).Err(); err != nil {
		return common.NewServerErr(err)
	}

	return nil
}
