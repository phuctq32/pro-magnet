package routesv1

import (
	"github.com/gin-gonic/gin"
	"pro-magnet/components/appcontext"
	ggoauth2 "pro-magnet/components/googleoauth2"
	hasher2 "pro-magnet/components/hasher"
	"pro-magnet/components/jwt"
	"pro-magnet/components/mailer/sendgrid"
	authrepo "pro-magnet/modules/auth/repository/redis"
	authapi "pro-magnet/modules/auth/transport/api"
	authuc "pro-magnet/modules/auth/usecase"
	userrepo "pro-magnet/modules/user/repository/mongo"
)

func NewAuthRouter(appCtx appcontext.AppContext, router *gin.RouterGroup) {
	// Setup dependencies
	userRepo := userrepo.NewUserRepository(appCtx.DBConnection())
	authRedisRepo := authrepo.NewAuthRedisRepository(appCtx.RedisClient())
	ggOauth := ggoauth2.NewGoogleOAuth2(
		appCtx.EnvConfigs().GoogleOAuth().ClientId(),
		appCtx.EnvConfigs().GoogleOAuth().ClientSecret(),
		appCtx.EnvConfigs().GoogleOAuth().RedirectUri(),
	)
	hasher := hasher2.NewBcryptHash(10)
	sgMailer := sendgrid.NewSendGridProvider(appCtx.EnvConfigs().Sendgrid().ApiKey())
	jwtProvider := jwt.NewJwtProvider()

	authUC := authuc.NewAuthUseCase(
		userRepo,
		authRedisRepo,
		ggOauth,
		hasher,
		sgMailer,
		appCtx.EnvConfigs().Sendgrid().FromEmail(),
		appCtx.EnvConfigs().Sendgrid().VerifyEmailTemplateId(),
		appCtx.EnvConfigs().Sendgrid().ResetPasswordEmailTemplateId(),
		appCtx.EnvConfigs().Sendgrid().VerificationURL(),
		appCtx.EnvConfigs().Sendgrid().ResetPasswordURL(),
		jwtProvider,
		appCtx.EnvConfigs().App().AccessSecret(),
		appCtx.EnvConfigs().App().RefreshSecret(),
		appCtx.EnvConfigs().App().AccessTokenExpiry(),
		appCtx.EnvConfigs().App().RefreshTokenExpiry(),
	)

	authHdl := authapi.NewAuthHandler(authUC)

	// Setup routes
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", authHdl.Register(appCtx))
		authRouter.POST("/verify", authHdl.Verify(appCtx))
		authRouter.POST("/send-verification-email", authHdl.SendVerificationEmail(appCtx))
		authRouter.POST("/login", authHdl.Login(appCtx))
		authRouter.POST("/refresh", authHdl.Refresh(appCtx))
		authRouter.POST("/google-login", authHdl.LoginWithGoogle(appCtx))
		authRouter.POST("/forgot-password", authHdl.ForgotPassword(appCtx))
		authRouter.POST("/reset-password", authHdl.ResetPassword(appCtx))
	}
}
