package routesv1

import (
	"github.com/gin-gonic/gin"
	"pro-magnet/components/appcontext"
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
	hasher := hasher2.NewBcryptHash(10)
	sgMailer := sendgrid.NewSendGridProvider(appCtx.EnvConfigs().Sendgrid().ApiKey())
	jwtProvider := jwt.NewJwtProvider()

	authUC := authuc.NewAuthUseCase(
		userRepo,
		authRedisRepo,
		hasher,
		sgMailer,
		appCtx.EnvConfigs().Sendgrid().FromEmail(),
		appCtx.EnvConfigs().Sendgrid().VerifyEmailTemplateId(),
		appCtx.EnvConfigs().Sendgrid().VerificationURL(),
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
	}
}
