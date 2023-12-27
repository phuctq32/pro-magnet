package authuc

import (
	"context"
	ggoauth2 "pro-magnet/components/googleoauth2"
	"pro-magnet/components/hasher"
	"pro-magnet/components/jwt"
	"pro-magnet/components/mailer"
	usermodel "pro-magnet/modules/user/model"
	wsmodel "pro-magnet/modules/workspace/model"
	wsmembermodel "pro-magnet/modules/workspacemember/model"
)

type WorkspaceRepository interface {
	Create(ctx context.Context, data *wsmodel.WorkspaceCreation) (*wsmodel.Workspace, error)
}

type WorkspaceMemberRepository interface {
	CreateMany(ctx context.Context, data *wsmembermodel.WorkspaceMembersCreate) error
}

type UserRepository interface {
	Create(ctx context.Context, data *usermodel.User) (*string, error)
	UserExist(ctx context.Context, email string) (bool, error)
	SetEmailVerified(ctx context.Context, id string) error
	FindByEmail(ctx context.Context, email string) (*usermodel.User, error)
	FindById(ctx context.Context, id string) (*usermodel.User, error)
	UpdatePasswordByEmail(ctx context.Context, email string, password string) error
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}

type AuthRedisRepository interface {
	SetVerifiedToken(ctx context.Context, token, userId string) error
	GetVerifiedUserId(ctx context.Context, verifiedToken string) (*string, error)
	SetResetToken(ctx context.Context, token, email string) error
	GetResetPasswordEmail(ctx context.Context, resetToken string) (*string, error)
}

type authUseCase struct {
	userRepo                     UserRepository
	redisRepo                    AuthRedisRepository
	wsRepo                       WorkspaceRepository
	wsMemberRepo                 WorkspaceMemberRepository
	ggOauth                      ggoauth2.GoogleOAuth
	hasher                       hasher.Hasher
	mailer                       mailer.Mailer
	fromEmail                    string
	verifyEmailTemplateId        string
	resetPasswordEmailTemplateId string
	verificationUrl              string
	resetPasswordUrl             string
	tokenProvider                jwt.TokenProvider
	atSecret                     string
	rtSecret                     string
	atExpiry                     int
	rtExpiry                     int
}

func NewAuthUseCase(
	userRepo UserRepository,
	authRedisRepo AuthRedisRepository,
	wsRepo WorkspaceRepository,
	wsMemberRepo WorkspaceMemberRepository,
	ggOauth ggoauth2.GoogleOAuth,
	hasher hasher.Hasher,
	mailer mailer.Mailer,
	fromEmail string,
	verifyEmailTemplateId string,
	resetPasswordEmailTemplateId string,
	verificationUrl string,
	resetPasswordUrl string,
	tokenProvider jwt.TokenProvider,
	atSecret string,
	rtSecret string,
	atExpiry int,
	rtExpiry int,
) *authUseCase {
	return &authUseCase{
		userRepo:                     userRepo,
		redisRepo:                    authRedisRepo,
		wsRepo:                       wsRepo,
		wsMemberRepo:                 wsMemberRepo,
		ggOauth:                      ggOauth,
		hasher:                       hasher,
		mailer:                       mailer,
		fromEmail:                    fromEmail,
		verifyEmailTemplateId:        verifyEmailTemplateId,
		resetPasswordEmailTemplateId: resetPasswordEmailTemplateId,
		verificationUrl:              verificationUrl,
		resetPasswordUrl:             resetPasswordUrl,
		tokenProvider:                tokenProvider,
		atSecret:                     atSecret,
		rtSecret:                     rtSecret,
		atExpiry:                     atExpiry,
		rtExpiry:                     rtExpiry,
	}
}
