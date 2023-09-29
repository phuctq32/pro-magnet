package authuc

import (
	"context"
	"pro-magnet/components/hasher"
	"pro-magnet/components/mailer"
	usermodel "pro-magnet/modules/user/model"
)

type UserRepository interface {
	Create(ctx context.Context, data *usermodel.User) (*string, error)
	CheckEmailExists(ctx context.Context, email string) error
	SetEmailVerified(ctx context.Context, id string) error
	FindByEmail(ctx context.Context, email string) (*usermodel.User, error)
}

type AuthRedisRepository interface {
	SetVerifiedToken(ctx context.Context, token, userId string) error
	GetVerifiedUserId(ctx context.Context, verifiedToken string) (*string, error)
}

type authUseCase struct {
	userRepo  UserRepository
	redisRepo AuthRedisRepository
	hasher    hasher.Hasher
	mailer    mailer.Mailer
}

func NewAuthUseCase(
	userRepo UserRepository,
	authRedisRepo AuthRedisRepository,
	hasher hasher.Hasher,
	mailer mailer.Mailer,
) *authUseCase {
	return &authUseCase{
		userRepo:  userRepo,
		redisRepo: authRedisRepo,
		hasher:    hasher,
		mailer:    mailer,
	}
}
