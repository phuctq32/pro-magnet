package authrepo

import "github.com/redis/go-redis/v9"

type authRedisRepository struct {
	cli *redis.Client
}

func NewAuthRedisRepository(cli *redis.Client) *authRedisRepository {
	return &authRedisRepository{cli: cli}
}
