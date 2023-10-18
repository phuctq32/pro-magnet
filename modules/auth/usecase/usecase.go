package authuc

import (
	"context"
	"pro-magnet/components/hasher"
	"pro-magnet/components/jwt"
	"pro-magnet/components/mailer"
	usermodel "pro-magnet/modules/user/model"
)

type UserRepository interface {
	Create(ctx context.Context, data *usermodel.User) (*string, error)
	CheckEmailExists(ctx context.Context, email string) error
	SetEmailVerified(ctx context.Context, id string) error
	FindByEmail(ctx context.Context, email string) (*usermodel.User, error)
	FindById(ctx context.Context, id string) (*usermodel.User, error)
}

type AuthRedisRepository interface {
	SetVerifiedToken(ctx context.Context, token, userId string) error
	GetVerifiedUserId(ctx context.Context, verifiedToken string) (*string, error)
}

type authUseCase struct {
	userRepo              UserRepository
	redisRepo             AuthRedisRepository
	hasher                hasher.Hasher
	mailer                mailer.Mailer
	fromEmail             string
	verifyEmailTemplateId string
	verificationUrl       string
	tokenProvider         jwt.TokenProvider
	atSecret              string
	rtSecret              string
	atExpiry              int
	rtExpiry              int
}

func NewAuthUseCase(
	userRepo UserRepository,
	authRedisRepo AuthRedisRepository,
	hasher hasher.Hasher,
	mailer mailer.Mailer,
	fromEmail string,
	verifyEmailTemplateId string,
	verificationUrl string,
	tokenProvider jwt.TokenProvider,
	atSecret string,
	rtSecret string,
	atExpiry int,
	rtExpiry int,
) *authUseCase {
	return &authUseCase{
		userRepo:              userRepo,
		redisRepo:             authRedisRepo,
		hasher:                hasher,
		mailer:                mailer,
		fromEmail:             fromEmail,
		verifyEmailTemplateId: verifyEmailTemplateId,
		verificationUrl:       verificationUrl,
		tokenProvider:         tokenProvider,
		atSecret:              atSecret,
		rtSecret:              rtSecret,
		atExpiry:              atExpiry,
		rtExpiry:              rtExpiry,
	}
}
