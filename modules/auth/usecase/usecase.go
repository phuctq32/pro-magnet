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
}

type AuthRedisRepository interface {
	SetVerifiedToken(ctx context.Context, token, userId string) error
}

type authUseCase struct {
	userRepo      UserRepository
	authRedisRepo AuthRedisRepository
	hasher        hasher.Hasher
	mailer        mailer.Mailer
}

func NewAuthUseCase(
	userRepo UserRepository,
	authRedisRepo AuthRedisRepository,
	hasher hasher.Hasher,
	mailer mailer.Mailer,
) *authUseCase {
	return &authUseCase{
		userRepo:      userRepo,
		authRedisRepo: authRedisRepo,
		hasher:        hasher,
		mailer:        mailer,
	}
}
